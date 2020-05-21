package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/projmanserver/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Initialize jwtKey
var jwtKey = []byte("my_secret_key")

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
var dbClient = dynamodb.New(sess)

// CreateUser will put a user into the database. Called when user signs up.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Create user struct and decode request into it
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	// Generate Hash
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error in hashing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// User password is now hashed
	user.UserID = uuid.New().String()
	user.Password = string(pass)
	log.Println(user)

	// Marshal the user into an item that can be stored into DynamoDB
	uv, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Println("Got error marshalling new movie item:")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Use put parameters in a way the DynamoDB sdk understands
	tableName := "Users"
	input := &dynamodb.PutItemInput{
		Item:      uv,
		TableName: aws.String(tableName),
	}

	// Put item into table
	_, err = dbClient.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Login logs user in
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Could not decode user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare user email and password to expected email and password

	// Set token parameters
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	tk := &models.Token{
		UserID: user.UserID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, tknStringError := token.SignedString(jwtKey)
	if tknStringError != nil {
		log.Println("Error in generating token string")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user
	json.NewEncoder(w).Encode(resp)
}

// JwtVerify verifies if client is authenticated or not
func JwtVerify(w http.ResponseWriter, r *http.Request) {

	var header = r.Header.Get("x-access-token") //Grab the token from the header

	header = strings.TrimSpace(header)

	if header == "" {
		//Token is missing, returns with error code 403 Unauthorized
		log.Println("Missing token header")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	tk := &models.Token{}

	tkn, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.Exception{Message: err.Error()})
		return
	}
	if tkn.Valid {
		log.Println("Token isn't valid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", tk.Email)))
}
