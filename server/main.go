package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/liuyh73/GraphQL-Service"
	"github.com/liuyh73/GraphQL-Service/server/service"
)

const defaultPort = "8080"

func NewServer() *negroni.Negroni {
	router := mux.NewRouter()
	initRoutes(router)

	n := negroni.Classic()

	n.UseHandler(router)
	return n
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/login", service.LoginHandler).Methods("POST")
	// router.Use(service.TokenMiddleware)
	// router.HandleFunc("/", service.ApiHandler).Methods("GET")
	router.HandleFunc("/", handler.Playground("GraphQL playground", "/query"))
	router.HandleFunc("/query", handler.GraphQL(GraphQL_Service.NewExecutableSchema(GraphQL_Service.Config{Resolvers: &GraphQL_Service.Resolver{}})))
	router.HandleFunc("/logout", service.LogoutHandler).Methods("POST", "GET")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := NewServer()
	server.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
