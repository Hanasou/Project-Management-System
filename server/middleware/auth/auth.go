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

	userID := uuid.New().String()
	user.UserID = userID
	tableName := "Users"

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
	log.Println("Put Successful")

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

	log.Println("User login: ", user)

	// Get item from table
	tableName := "Users"
	userEmail := user.Email
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(userEmail),
			},
		},
	}
	getResult, _ := dbClient.GetItem(getItemInput)
	uv := models.User{}
	dynamodbattribute.UnmarshalMap(getResult.Item, &uv)
	log.Println("User Retrieved: ", uv)

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		var header = r.Header.Get("Authorization") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			msg := "Missing token header"
			log.Println(msg)
			http.Error(w, msg, http.StatusForbidden)
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			msg := "Could not parse token"
			log.Println(msg)
			http.Error(w, msg, http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
