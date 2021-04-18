package main

import (
	"gatewayCore/train/graph"
	"gatewayCore/train/graph/generated"
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

	http.Handle("/train/playground", playground.Handler("GraphQL playground", "/train"))
	http.Handle("/train", srv)

	log.Printf("connect to http://localhost:%s/train/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
