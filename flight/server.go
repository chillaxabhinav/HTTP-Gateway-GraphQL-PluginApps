package main

import (
	"gatewayCore/flight/graph"
	"gatewayCore/flight/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/flight/playground", playground.Handler("GraphQL playground", "/flight"))
	http.Handle("/flight", srv)

	log.Printf("connect to http://localhost:%s/flight/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
