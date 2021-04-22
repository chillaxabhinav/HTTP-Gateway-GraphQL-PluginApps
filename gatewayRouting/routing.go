package gatewayRouting

import (
	serviceLinking "gatewayCore/serviceLinking"
	utils "gatewayCore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GraphQLModuleEndpointMapping map[string]serviceLinking.ModuleInfoMap

func RoutingToGQL(c *gin.Context) {

	mapping := GraphQLModuleEndpointMapping{
		"flight": serviceLinking.ModuleInfoMap{
			Internal:   true,
			Endpoint:   "/flight",
			ModuleName: "flight",
		},
		"train": serviceLinking.ModuleInfoMap{
			Internal:   true,
			Endpoint:   "/train",
			ModuleName: "train",
		},
		"search": serviceLinking.ModuleInfoMap{
			Internal:   false,
			Endpoint:   "http://localhost:8000/search",
			ModuleName: "search",
		},
	}

	incomingModuleRequest := c.GetHeader("service")

	if incomingModuleRequest == "" {
		c.Data(
			http.StatusBadGateway,
			"application/json",
			utils.GetError("no nodule header defined in query, expects a module header", http.StatusNotFound))
		c.Abort()
		return
	}

	moduleInfo, ok := mapping[incomingModuleRequest]

	if !ok {
		c.Data(
			http.StatusNotFound,
			"application/json",
			utils.GetError("service not found", http.StatusNotFound))
		return
	}

	serviceLinking.GraphQLHandlerOneGatewayMiddleware(c, moduleInfo)

}
