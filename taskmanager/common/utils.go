package common

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

type (
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus int `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = Configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}

	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; chartset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}