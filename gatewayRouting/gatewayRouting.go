package gatewayRouting

import (
	serviceLinking "gatewayCore/gatewayServiceLinking"

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
		c.JSON(422, gin.H{"error": "No nodule header defined in query, expects a module header", "data": nil})
		c.Abort()
		return
	}

	moduleEndpoint, ok := mapping[incomingModuleRequest]

	if !ok {
		c.JSON(422, gin.H{"error": "Module not found", "data": nil})
		c.Abort()
		return
	}

	serviceLinking.GraphQLHandlerOneGatewayMiddleware(c, moduleEndpoint)

}
