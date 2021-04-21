package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"

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

func reverseProxyExternalServices(c *gin.Context, target string) {
	remote, err := url.Parse(target)

	if err != nil {
		c.JSON(422, gin.H{"error": "Cannot proxt to target: " + target, "data": nil})
		c.Abort()
		return
	}

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path
	}

	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func GraphQLHandler(endpoint string) gin.HandlerFunc {
	if endpoint == "/flight" {
		srv := handler.NewDefaultServer(flightGenerated.NewExecutableSchema(flightGenerated.Config{Resolvers: &flightGraph.Resolver{}}))
		return func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		}
	} else if endpoint == "/train" {
		srv := handler.NewDefaultServer(trainGenerated.NewExecutableSchema(trainGenerated.Config{Resolvers: &trainGraph.Resolver{}}))
		return func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		}
	} else {
		return func(c *gin.Context) {
			c.JSON(422, gin.H{"error": "No endpoint", "data": nil})
			c.Abort()
		}
	}
}

func GraphQLHandlerOneGatewayMiddleware(c *gin.Context, endpoint string) {
	if endpoint == "/flight" {
		srv := handler.NewDefaultServer(flightGenerated.NewExecutableSchema(flightGenerated.Config{Resolvers: &flightGraph.Resolver{}}))
		srv.ServeHTTP(c.Writer, c.Request)
		return
	} else if endpoint == "/train" {
		srv := handler.NewDefaultServer(trainGenerated.NewExecutableSchema(trainGenerated.Config{Resolvers: &trainGraph.Resolver{}}))
		srv.ServeHTTP(c.Writer, c.Request)
		return
	} else {
		reverseProxyExternalServices(c, endpoint)
		return
	}
}
