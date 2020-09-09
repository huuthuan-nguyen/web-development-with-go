package main

import (
	"log"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/routers"
)

// entry point of program
func main() {
	
	// call startup logic
	common.StartUp()

	// get mux router object
	router := routers.InitRoutes()

	// create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}