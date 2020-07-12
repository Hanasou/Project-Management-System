package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
	"github.com/projmanserver/models"
)

type body struct {
	Key1 string
	Key2 string
}

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var dbClient = dynamodb.New(sess)

// TestFunc test functionality of server
func TestFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from "+r.URL.Path[1:])
}

// TestPathParams tests getting path parameters
func TestPathParams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("Getting path parameter")
	params := mux.Vars(r)
	id := params["id"]
	log.Println(id)
	json.NewEncoder(w).Encode(id)
}

// TestPost is a dummy post request
func TestPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("starting test post request")
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	log.Println(user)

	tableName := "Users"
	indexName := "Email-index"
	userEmail := user.Email
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("#em = :email"),
		ExpressionAttributeNames: map[string]*string{
			"#em": aws.String("Email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(userEmail),
			},
		},
	}
	queryResult, _ := dbClient.Query(queryInput)
	userList := []models.User{}
	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &userList)
	log.Println("User Retrieved: ", userList)
	json.NewEncoder(w).Encode(userList)
}

// ReadRequest func reads from request body
func ReadRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := body{}
	json.NewDecoder(r.Body).Decode(&b)
	json.NewEncoder(w).Encode(b)
}

// AuthRequest only works if user is authenticated
func AuthRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := body{}
	json.NewDecoder(r.Body).Decode(&b)
	json.NewEncoder(w).Encode(b)
}
