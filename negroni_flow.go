package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
)

func middlewareFirst(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Middleware First - Before Handler")
	next(w, r)
	log.Println("Middleware First - After Handler")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Middleware Second - Before Handler")
	if r.URL.Path == "/message" {
		if r.URL.Query().Get("password") == "pass123" {
			log.Println("Authorized to the system")
			next(w, r)
		} else {
			log.Println("Failed to authorize to the system")
			return
		}
	} else {
		next(w, r)
	}
	log.Println("Middleware Second - After Handler")
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index Handler")
	fmt.Fprintf(w, "Welcome")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing message Handler")
	fmt.Fprintf(w, "HTTP Middlware is awesome")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.icon", iconHandler)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/message", message)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middlewareFirst))
	n.Use(negroni.HandlerFunc(middlewareSecond))
	n.UseHandler(mux)
	n.Run(":8080")
}