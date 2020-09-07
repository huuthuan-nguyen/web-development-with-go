package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/context"
	"github.com/codegangsta/negroni"
)

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("X-AppToken")
	if token == "bXlVc2VybmFtZTpteVBhc3N3b3Jk" {
		log.Printf("Authorized to the system")
		context.Set(r, "user", "Thuan Nguyen")
		next(w, r)
	} else {
		http.Error(w, "Not Authroized", 401)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	fmt.Fprintf(w, "Welcome %s!", user)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(Authorize))
	n.UseHandler(mux)
	n.Run(":8080")
}