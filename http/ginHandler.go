package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	trainGenerated "gatewayCore/train/graph/generated"

	trainGraph "gatewayCore/train/graph"

	flightGenerated "gatewayCore/flight/graph/generated"

	flightGraph "gatewayCore/flight/graph"
)

func PlaygroundHandler(endPoint string) gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", endPoint)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GraphQLHandler(endpoint string) gin.HandlerFunc {
	if endpoint == "/flight" {
		srv := handler.NewDefaultServer(flightGenerated.NewExecutableSchema(flightGenerated.Config{Resolvers: &flightGraph.Resolver{}}))
		return func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		}
	}
	srv := handler.NewDefaultServer(trainGenerated.NewExecutableSchema(trainGenerated.Config{Resolvers: &trainGraph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
