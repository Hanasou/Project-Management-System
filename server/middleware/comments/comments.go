package comments

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/projmanserver/models"
)

const tableName = "Comments"

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var dbClient = dynamodb.New(sess)

// AddComment adds a comment to an issue
func AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight request
	log.Println("Request Method", r.Method)
	if r.Method != "POST" {
		return
	}

	requestObject := models.Comment{}
	json.NewDecoder(r.Body).Decode(&requestObject)

	commentID := uuid.New().String()
	requestObject.CommentID = commentID

	now := time.Now()
	nowText, _ := now.MarshalText()
	requestObject.Timestamp = string(nowText)

	cav, err := dynamodbattribute.MarshalMap(requestObject)
	if err != nil {
		msg := "Could not marshal comment"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      cav,
		TableName: aws.String(tableName),
	}

	dbClient.PutItem(input)

	json.NewEncoder(w).Encode(requestObject)
}

// GetComments displays all the comments for a particular issue
func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting comments")

	// Get the IssueID from the path parameters
	issueID := mux.Vars(r)["issueID"]
	log.Println(issueID)

	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
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

	// Calculate time since comment was posted for display
	// Kind of confusing here since we're updating the Timestamp field
	for _, c := range comments {
		parsedTime, err := time.Parse(time.RFC3339, c.Timestamp)
		if err != nil {
			msg := "Cannot parse time"
			log.Println(msg)
			http.Error(w, msg, http.StatusBadRequest)
		}
		c.Timestamp = time.Since(parsedTime).String()
	}

	log.Println(comments)
	json.NewEncoder(w).Encode(comments)
}
