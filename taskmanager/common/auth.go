package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

// using asymetric crypto/RSA key
const (
	PRIVATE_KEY_PATH = "keys/app.rsa"
	PUBLIC_KEY_PATH = "keys/app.rsa.pub"
)

// private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read the key files before starting http handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(PRIVATE_KEY_PATH)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
	verifyKey, err = ioutil.ReadFile(PUBLIC_KEY_PATH)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set claims for JWT token
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// Set the expire time for JWT token
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Middleware for validating JWT Tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HanlderFunc) {
	// validate the token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		// verify the token with public key, which is the counter part of private key
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: // JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return

			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Access Token",
					500,
				)
				return
			}

		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing Access Token!",
				500,
			)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}