package projects

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/projmanserver/models"
)

const tableName = "Projects"

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var dbClient = dynamodb.New(sess)

// AddProject adds a project to the projects table
func AddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Unmarshal request body into project object
	project := models.Project{}
	json.NewDecoder(r.Body).Decode(&project)
	projectID := uuid.New().String()
	project.ProjectID = projectID

	log.Println("Project request", project)
	// Email should be sent inside path parameter
	userEmail := mux.Vars(r)["email"]
	project.UserEmail = userEmail

	pv, err := dynamodbattribute.MarshalMap(project)
	if err != nil {
		msg := "Could not unmarshal"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Use put parameters in a way the DynamoDB sdk understands
	input := &dynamodb.PutItemInput{
		Item:      pv,
		TableName: aws.String(tableName),
	}

	// Put item into table
	dbClient.PutItem(input)

	json.NewEncoder(w).Encode(project)
}

// GetProjects gets all projects with a given email
func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting Projects...")
	// Email should be sent inside path parameter
	userEmail := mux.Vars(r)["email"]
	log.Println(userEmail)

	// Query for items that match this userEmail
	indexName := "Email-index"
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("#em = :email"),
		ExpressionAttributeNames: map[string]*string{
			"#em": aws.String("UserEmail"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(userEmail),
			},
		},
	}

	queryResult, _ := dbClient.Query(queryInput)
	var projects = []models.Project{}

	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &projects)

	log.Println(projects)
	json.NewEncoder(w).Encode(projects)
}
