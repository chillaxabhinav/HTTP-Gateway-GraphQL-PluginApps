package main

import (
	serviceLinking "gatewayCore/gatewayServiceLinking"

	gatewayRouting "gatewayCore/gatewayRouting"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {

	server := gin.Default()

	// ===== Train GraphQL Module ===== //

	server.GET("/train/playground", serviceLinking.PlaygroundHandler("/train"))

	server.POST("/train", serviceLinking.GraphQLHandler("/train"))

	// ================================ //

	// ===== Flight GraphQL Module ==== //

	server.GET("/flight/playground", serviceLinking.PlaygroundHandler("/flight"))

	server.POST("/flight", serviceLinking.GraphQLHandler("/flight"))

	// ================================= //

	// ===== One Gateway ======= //

	server.POST("/gateway", gatewayRouting.RoutingToGQL)

	// ========================= //

	server.Run(defaultPort)

}
