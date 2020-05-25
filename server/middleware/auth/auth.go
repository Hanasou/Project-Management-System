package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/gddo/httputil/header"
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

// TestPost is a dummy post request
func TestPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Println("starting test post request")

	// Create user struct and decode request into it
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	userEmail := user.Email
	log.Println("User Email: " + userEmail)

	// Check if email exists in the database
	tableName := "Users"
	result, _ := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(userEmail),
			},
		},
	})
	log.Println(result)
	if result.Item != nil {
		msg := "User already exists"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Generate Hash
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		msg := "Error in hashing"
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	// User password is now hashed
	user.Password = string(pass)
	log.Println(user)

	// Marshal the user into an item that can be stored into DynamoDB
	uv, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		msg := "Could not unmarshal"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Use put parameters in a way the DynamoDB sdk understands
	input := &dynamodb.PutItemInput{
		Item:      uv,
		TableName: aws.String(tableName),
	}

	// Put item into table
	dbClient.PutItem(input)

	json.NewEncoder(w).Encode(user)
}

// CreateUser will put a user into the database. Called when user signs up.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			log.Println(msg)
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	// Create user struct and decode request into it
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	userEmail := user.Email
	log.Println("User object: " + user.Email)

	// Check if email exists in the database
	tableName := "Users"
	result, _ := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(userEmail),
			},
		},
	})
	log.Println(result)
	if result.Item != nil {
		msg := "User already exists"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Generate Hash
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		msg := "Error in hashing"
		log.Println(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	// User password is now hashed
	user.Password = string(pass)
	log.Println(user)

	// Marshal the user into an item that can be stored into DynamoDB
	uv, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		msg := "Could not unmarshal"
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Use put parameters in a way the DynamoDB sdk understands
	input := &dynamodb.PutItemInput{
		Item:      uv,
		TableName: aws.String(tableName),
	}

	// Put item into table
	dbClient.PutItem(input)

	json.NewEncoder(w).Encode(user)
}

// Login logs user in
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Unmarshal the request body into user struct
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	// Get item from table
	tableName := "Users"
	userEmail := user.Email
	result, _ := dbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(userEmail),
			},
		},
	})

	// uv is the item from the database that corresponds to the user input
	// Unmarshal the result into uv
	uv := models.User{}
	err := dynamodbattribute.UnmarshalMap(result.Item, &uv)
	if err != nil {
		log.Println("Could not unmarshal record")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Compare password inside request to password inside database
	errf := bcrypt.CompareHashAndPassword([]byte(uv.Password), []byte(user.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid Password"))
		return
	}

	// Set token parameters
	tkExpiresAt := time.Now().Add(time.Minute * 100000).Unix()
	rtExpiresAt := time.Now().Add(time.Hour * 24).Unix()
	tk := &models.Token{
		Email: uv.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: tkExpiresAt,
		},
	}
	rt := &models.Token{
		Email: uv.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: rtExpiresAt,
		},
	}

	// Set token string and return it to client
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println("Error in generating access token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rt)
	rtString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		log.Println("Error in generating refresh token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["refreshToken"] = rtString
	resp["userEmail"] = user.Email
	resp["expiresIn"] = tkExpiresAt
	json.NewEncoder(w).Encode(resp)
}

// JwtVerify is middleware that verifies if client is authenticated or not
func JwtVerify(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			log.Println("Missing token header")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Exception{Message: "Not Authorized"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
