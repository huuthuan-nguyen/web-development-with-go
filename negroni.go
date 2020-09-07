package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}