package issues

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/projmanserver/models"
)

const tableName = "Issues"

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var dbClient = dynamodb.New(sess)

// CreateIssue creates a new issue for a project
func CreateIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Creating new issue")

	log.Println("Request Method", r.Method)
	if r.Method != "POST" {
		return
	}

	// Unmarshal request object into issue struct
	issue := models.Issue{}
	json.NewDecoder(r.Body).Decode(&issue)
	issueID := uuid.New().String()
	issue.IssueID = issueID
	issue.Status = "In Progress"

	now := time.Now()
	nowText, err := now.MarshalText()
	if err != nil {
		msg := "Could not marshal text"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	issue.Timestamp = string(nowText)

	// Marshal issue into dynamodb attribute value
	iav, err := dynamodbattribute.MarshalMap(issue)
	if err != nil {
		msg := "Could not marshal into attribute value"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Use put parameters in a way the DynamoDB sdk understands
	input := &dynamodb.PutItemInput{
		Item:      iav,
		TableName: aws.String(tableName),
	}

	// Put item into table
	dbClient.PutItem(input)

	json.NewEncoder(w).Encode(issue)
}

// GetIssues gets all the issues for a particular project
func GetIssues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting issues")

	// Get projectID from path parameter
	projectID := mux.Vars(r)["projectID"]
	log.Println(projectID)

	// Query for items that match this userEmail
	indexName := "ProjectID-index"
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("#pid = :pid"),
		ExpressionAttributeNames: map[string]*string{
			"#pid": aws.String("ProjectID"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pid": {
				S: aws.String(projectID),
			},
		},
	}

	log.Println(queryInput.ExpressionAttributeValues)
	queryResult, _ := dbClient.Query(queryInput)
	issues := []models.Issue{}

	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &issues)

	json.NewEncoder(w).Encode(issues)
}

// GetIssue gets a single issue
func GetIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting issue")

	// Project ID should be in path parameter
	projectID := mux.Vars(r)["projectID"]
	issueID := mux.Vars(r)["issueID"]
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"IssueID": {
				S: aws.String(issueID),
			},
			"ProjectID": {
				S: aws.String(projectID),
			},
		},
	}
	issue := models.Issue{}
	getResult, err := dbClient.GetItem(getItemInput)
	if err != nil {
		msg := "Could not get items"
		log.Println(msg)
		log.Println(err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	err = dynamodbattribute.UnmarshalMap(getResult.Item, &issue)
	if err != nil {
		msg := "Could not unmarshal"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(issue)
}

// UpdateIssue updates an issue
// Request Body example template
// {IssueID, ProjectID, Title, Description, Type, Priority, Status}
func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// handle preflight request
	if r.Method != "POST" {
		return
	}

	updateRequest := models.Issue{}
	json.NewDecoder(r.Body).Decode(&updateRequest)

	uAttributeValues := map[string]*dynamodb.AttributeValue{}
	uAttributeNames := map[string]*string{}
	reqFields := reflect.TypeOf(updateRequest)
	reqValues := reflect.ValueOf(updateRequest)
	var expressionBuilder strings.Builder
	expressionBuilder.WriteString("SET ")

	// Build ExpressionAttribute Values
	// And also Update expression?
	// Update expressions are formatted like "Operation FieldName = AttributeValue"
	// e.g. "SET Title = :ti, Description = :d"
	for i := 0; i < reqFields.NumField(); i++ {
		field := reqFields.Field(i).Name
		value := reqValues.Field(i).String()
		if value == "" {
			continue
		}
		nameAttribute := &field
		itemAttribute := &dynamodb.AttributeValue{
			S: aws.String(value),
		}
		wildcard := ""
		fieldWildcard := ""
		switch field {
		case "Title":
			fieldWildcard = "#ti"
			wildcard = ":ti"
		case "Description":
			fieldWildcard = "#d"
			wildcard = ":d"
		case "Status":
			fieldWildcard = "#s"
			wildcard = ":s"
		case "Type":
			fieldWildcard = "#ty"
			wildcard = ":ty"
		case "Priority":
			fieldWildcard = "#p"
			wildcard = ":p"
		}

		if wildcard != "" {
			expressionBuilder.WriteString(fieldWildcard + " = " + wildcard + ", ")
			uAttributeNames[fieldWildcard] = nameAttribute
			uAttributeValues[wildcard] = itemAttribute
		}
	}

	// Convert to string and slice off the trailing comma
	s := expressionBuilder.String()
	updateExpression := s[:len(s)-2]

	updateItemInput := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames:  uAttributeNames,
		ExpressionAttributeValues: uAttributeValues,
		TableName:                 aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"IssueID": {
				S: aws.String(updateRequest.IssueID),
			},
			"ProjectID": {
				S: aws.String(updateRequest.ProjectID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(updateExpression),
	}

	_, err := dbClient.UpdateItem(updateItemInput)
	if err != nil {
		msg := updateExpression
		log.Println(msg)
		log.Println(err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Get back updated item and send it back to user
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"IssueID": {
				S: aws.String(updateRequest.IssueID),
			},
			"ProjectID": {
				S: aws.String(updateRequest.ProjectID),
			},
		},
	}

	resIssue := models.Issue{}
	getResult, err := dbClient.GetItem(getItemInput)
	if err != nil {
		msg := "Could not get item"
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	err = dynamodbattribute.UnmarshalMap(getResult.Item, &resIssue)
	if err != nil {
		msg := "Could not unmarshal"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(resIssue)
}

// DeleteIssue deletes an issue
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Deleting Issue")

	// Handle preflight request
	if r.Method != "DELETE" {
		return
	}
}
