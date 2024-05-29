package routes

import (
	"github.com/Cyber-Canum/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Route) {
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
}
