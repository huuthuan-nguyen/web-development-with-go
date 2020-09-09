package routers

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// routes for the User entity
	router = SetUserRoutes(router)
	// routes for the Task entity
	router = SetTaskRoutes(router)
	// routes for the Note entity
	router = SetNoteRoutes(router)

	return router
}