package serviceLinking

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	utils "gatewayCore/utils"

	trainGenerated "gatewayCore/train/graph/generated"

	trainGraph "gatewayCore/train/graph"

	flightGenerated "gatewayCore/flight/graph/generated"

	flightGraph "gatewayCore/flight/graph"
)

type ModuleInfoMap struct {
	Internal   bool
	Endpoint   string
	ModuleName string
}

func reverseProxyExternalServices(c *gin.Context, endpoint string, moduleName string) {
	remote, err := url.Parse(endpoint)

	if err != nil {
		c.Data(
			http.StatusBadGateway,
			"application/json",
			utils.GetError("cannot proxy "+moduleName+" to endpoint: "+endpoint, http.StatusBadGateway))
		c.Abort()
		return
	}

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path
	}

	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, e error) {
		c.Data(
			http.StatusBadGateway,
			"application/json",
			utils.GetError("service unavailable", http.StatusBadGateway))
		c.Abort()
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

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
	} else if endpoint == "/train" {
		srv := handler.NewDefaultServer(trainGenerated.NewExecutableSchema(trainGenerated.Config{Resolvers: &trainGraph.Resolver{}}))
		return func(c *gin.Context) {
			srv.ServeHTTP(c.Writer, c.Request)
		}
	} else {
		return func(c *gin.Context) {
			c.JSON(422, gin.H{"error": "No endpoint", "data": nil, "code": 422})
			c.Abort()
		}
	}
}

func GraphQLHandlerOneGatewayMiddleware(c *gin.Context, moduleInfo ModuleInfoMap) {

	if moduleInfo.Internal {
		if moduleInfo.ModuleName == "flight" {
			srv := handler.NewDefaultServer(flightGenerated.NewExecutableSchema(flightGenerated.Config{Resolvers: &flightGraph.Resolver{}}))
			srv.ServeHTTP(c.Writer, c.Request)
			return
		} else if moduleInfo.ModuleName == "train" {
			srv := handler.NewDefaultServer(trainGenerated.NewExecutableSchema(trainGenerated.Config{Resolvers: &trainGraph.Resolver{}}))
			srv.ServeHTTP(c.Writer, c.Request)
			return
		}

	} else {
		reverseProxyExternalServices(c, moduleInfo.Endpoint, moduleInfo.ModuleName)
		return
	}
}
