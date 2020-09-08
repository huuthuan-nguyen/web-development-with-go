package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"crypto/rsa"
)

// using asymmetric crypto/RSA keys
// location of the files used for signing and verification
const (
	privKeyPath = "keys/app.rsa" // openssl genrsa -out app.rsa 1024
	pubKeyPath = "keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout -out app.rsa.pub
)

// verify key and sign key
var (
	verifyKeyByte, signKeyByte []byte
	signKey *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

// struct User for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// read the key files before starting http handlers
func init() {
	var err error

	signKeyByte, err := ioutil.ReadFile(privKeyPath)
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)
	if err != nil {
		log.Fatalf("Error reading private key: %s\n", err)
		return
	}

	verifyKeyByte, err = ioutil.ReadFile(pubKeyPath)
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
	if err != nil {
		log.Fatalf("Error reading public key: %s\n", err)
		return
	}
}

// rads the login credentials, checks them and creates JWT the token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	// decode into User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}

	// validate user credentials.
	if user.UserName != "Thuan" && user.Password != "pass" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}

	// create a signer for rsa 256
	// t := jwt.New(jwt.GetSigningMethod("RS256"))
	token := jwt.New(jwt.SigningMethodRS256)

	claims := make(jwt.MapClaims)
	// set our claims
	claims["iss"] = "admin"
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.UserName, "Member"}

	// set expired time
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error when signing token")
		log.Printf("Token signing error: %v\n", err)
		return
	}

	response := Token{tokenString}
	jsonResponse(response, w)
}

// only accessible with a valid token
func authHandler(w http.ResponseWriter, r *http.Request) {
	// validate the token
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		// since we only use one private key to sign the token,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // something was wrong during the validation
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token Expired, get a new one.")
				return

			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while parsing token!")
				log.Printf("VlidationError error: %+v\n", vErr.Errors)
				return
			}

		default: // something else went wrong
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while parsing token!")
			log.Printf("Token parse err: %v\n", err)
			return
		}
	}

	if token.Valid {
		response := Response{"Authorized to the system"}
		jsonResponse(response, w)
	} else {
		response := Response{"Invalid Token"}
		jsonResponse(response, w)
	}
}

type Response struct {
	Text string `json:"text"`
}

type Token struct {
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// entry point of program
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}