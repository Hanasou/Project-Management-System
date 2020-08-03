package issues

import (
	"encoding/json"
	"errors"
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

// GetIssuesByUser gets issues for a user
func GetIssuesByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting issues by user")

	// Get user email from path parameter
	userEmail := mux.Vars(r)["userEmail"]
	log.Println("User: ", userEmail)

	// Query for items that match this userEmail
	indexName := "Creator-index"
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("#em = :email"),
		ExpressionAttributeNames: map[string]*string{
			"#em": aws.String("Creator"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(userEmail),
			},
		},
	}

	// Run query
	queryResult, err := dbClient.Query(queryInput)
	if err != nil {
		msg := "Could not get issues"
		log.Println(msg)
		log.Println(err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Unmarshal items into slice of issues
	issues := []models.Issue{}
	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &issues)

	// Return items to client
	json.NewEncoder(w).Encode(issues)
}

// GetIssues gets all the issues for a particular project
func GetIssues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting issues by project")

	// Get projectID from path parameter
	projectID := mux.Vars(r)["projectID"]
	log.Println(projectID)

	// Query for items that match this Project ID
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
	queryResult, err := dbClient.Query(queryInput)
	if err != nil {
		msg := "Could not get issues"
		log.Println(msg)
		log.Println(err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
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
		log.Println(err.Error())
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
// This should delete an issue from the issue table and all its associated comments
// Request should take in IssueID, ProjectID, and a CommentID array
func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Deleting Issue")

	// Handle preflight request
	if r.Method != "DELETE" {
		return
	}

	type deleteResponse struct {
		IssueID   string
		ProjectID string
	}

	issueID := mux.Vars(r)["issueID"]
	projectID := mux.Vars(r)["projectID"]
	responseObject := deleteResponse{
		IssueID:   issueID,
		ProjectID: projectID,
	}

	err := RemoveIssue(responseObject.IssueID, responseObject.ProjectID)
	if err != nil {
		msg := "RemoveItem error"
		log.Println(err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&responseObject)
}

// RemoveIssue deletes an issue.
func RemoveIssue(issueID, projectID string) error {

	// Delete all the comments associated with the issue from the comments table
	//Query for the comments associated with an issue
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String("Comments"),
		KeyConditionExpression: aws.String("#iid = :iid"),
		ExpressionAttributeNames: map[string]*string{
			"#iid": aws.String("IssueID"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":iid": {
				S: aws.String(issueID),
			},
		},
	}
	queryResult, _ := dbClient.Query(queryInput)
	comments := []models.Comment{}

	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &comments)

	// Iterate over comments slice and delete them
	for _, v := range comments {
		deleteCommentInput := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"IssueID": {
					S: aws.String(v.IssueID),
				},
				"CommentID": {
					S: aws.String(v.CommentID),
				},
			},
			TableName: aws.String("Comments"),
		}
		_, err := dbClient.DeleteItem(deleteCommentInput)
		if err != nil {
			msg := "Could not call DeleteItem on comment"
			log.Println(msg)
			log.Println(err)
			return errors.New(msg)
		}
	}

	// Delete item from issue table
	deleteIssueInput := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"IssueID": {
				S: aws.String(issueID),
			},
			"ProjectID": {
				S: aws.String(projectID),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := dbClient.DeleteItem(deleteIssueInput)
	if err != nil {
		msg := "Could not call DeleteItem on issue"
		log.Println(msg)
		log.Println(err)
		return errors.New(msg)
	}

	return nil
}
