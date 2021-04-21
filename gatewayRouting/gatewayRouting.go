package gatewayRouting

import (
	http "gatewayCore/httpPluginLink"

	"github.com/gin-gonic/gin"
)

type GraphQLModuleEndpointMapping map[string]string

func RoutingToGQL(c *gin.Context) {

	mapping := GraphQLModuleEndpointMapping{
		"flight":   "/flight",
		"train":    "/train",
		"external": "http://localhost:8000/query",
	}

	incomingModuleRequest := c.GetHeader("module")

	if incomingModuleRequest == "" {
		c.JSON(422, gin.H{"error": "No Module defined in query", "data": nil})
		c.Abort()
		return
	}

	moduleEndpoint := mapping[incomingModuleRequest]

	http.GraphQLHandlerOneGatewayMiddleware(c, moduleEndpoint)

}
