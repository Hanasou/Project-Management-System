package teams

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
	"github.com/projmanserver/models"
)

const tableName = "Teams"

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var dbClient = dynamodb.New(sess)

// CreateTeam creates a team
// Request Body:
/**
{
	ProjectID string
	Members []
	Project string
}
*/
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Creating new team")
	log.Println("Request Method", r.Method)
	if r.Method != "POST" {
		return
	}

	type requestObject struct {
		ProjectID   string
		ProjectName string
		ProjectDesc string
		Members     []string
	}
	reqObject := requestObject{}
	json.NewDecoder(r.Body).Decode(&reqObject)

	// Take all the team members in the request list and put them into DynamoDB
	// Data is formatted like this for querying purposes
	for _, mem := range reqObject.Members {

		log.Println(mem)

		// Check if member exists in database
		getItemInput := &dynamodb.GetItemInput{
			TableName: aws.String("Users"),
			Key: map[string]*dynamodb.AttributeValue{
				"Email": {
					S: aws.String(mem),
				},
			},
		}
		memItem, _ := dbClient.GetItem(getItemInput)
		if memItem.Item == nil {
			msg := "User doesn't exist"
			log.Println(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		// Add team member to project table
		insertProject := models.Project{
			ProjectID:   reqObject.ProjectID,
			Title:       reqObject.ProjectName,
			Description: reqObject.ProjectDesc,
			Email:       mem,
		}
		err := addMemberToProject(insertProject)
		if err != nil {
			msg := "Insertion into project table failed"
			log.Println(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		teamMember := struct {
			ProjectID   string
			Member      string
			Project     string
			ProjectDesc string
		}{
			ProjectID:   reqObject.ProjectID,
			Member:      mem,
			Project:     reqObject.ProjectName,
			ProjectDesc: reqObject.ProjectDesc,
		}
		tav, err := dynamodbattribute.MarshalMap(teamMember)
		if err != nil {
			msg := "Could not marshal into attribute value"
			log.Println(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		input := &dynamodb.PutItemInput{
			Item:      tav,
			TableName: aws.String(tableName),
		}
		dbClient.PutItem(input)
	}

	json.NewEncoder(w).Encode(reqObject)
}

// GetTeams gets all the teams that you're part of
// Path parameters: user email
func GetTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	log.Println("Getting teams")

	// Member that is getting queried is inside path parameter
	memberEmail := mux.Vars(r)["memberEmail"]
	indexName := "Member-index"
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(indexName),
		KeyConditionExpression: aws.String("#mEmail = :mEmail"),
		ExpressionAttributeNames: map[string]*string{
			"#mEmail": aws.String("Member"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":mEmail": {
				S: aws.String(memberEmail),
			},
		},
	}

	queryResult, _ := dbClient.Query(queryInput)
	type dbTeam struct {
		ProjectID   string
		Member      string
		Project     string
		ProjectDesc string
	}
	dbTeams := []dbTeam{}
	outputTeams := []models.Team{}

	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &dbTeams)

	for _, t := range dbTeams {
		outputTeam := models.Team{}
		pid := t.ProjectID
		pDesc := t.ProjectDesc
		queryInput := &dynamodb.QueryInput{
			TableName:              aws.String(tableName),
			KeyConditionExpression: aws.String("#pid = :pid"),
			ExpressionAttributeNames: map[string]*string{
				"#pid": aws.String("ProjectID"),
			},
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":pid": {
					S: aws.String(pid),
				},
			},
		}
		queryResult, _ := dbClient.Query(queryInput)
		currTeams := []dbTeam{}
		dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &currTeams)
		outputTeam.ProjectID = pid
		outputTeam.Project = currTeams[0].Project
		outputTeam.ProjectDesc = pDesc
		for _, team := range currTeams {
			outputTeam.Members = append(outputTeam.Members, team.Member)
		}
		outputTeams = append(outputTeams, outputTeam)
	}

	json.NewEncoder(w).Encode(outputTeams)
}

// GetTeam gets a specific team
// Formats for teams are different so data transformation is required
// Path parameter: ProjectID
func GetTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	projectID := mux.Vars(r)["projectID"]
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
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
	queryResult, _ := dbClient.Query(queryInput)
	type dbTeam struct {
		ProjectID string
		Member    string
		Project   string
	}
	teams := []dbTeam{}
	dynamodbattribute.UnmarshalListOfMaps(queryResult.Items, &teams)
	log.Println(teams)

	respTeam := models.Team{
		ProjectID: projectID,
		Members:   []string{},
		Project:   teams[0].Project,
	}

	for _, t := range teams {
		respTeam.Members = append(respTeam.Members, t.Member)
	}

	json.NewEncoder(w).Encode(respTeam)
}

func addMemberToProject(project models.Project) error {
	pv, err := dynamodbattribute.MarshalMap(project)
	if err != nil {
		msg := "Could not unmarshal"
		log.Println(msg)
		return errors.New(msg)
	}

	input := &dynamodb.PutItemInput{
		Item:      pv,
		TableName: aws.String("Projects"),
	}

	dbClient.PutItem(input)
	return nil
}
