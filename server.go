package main

import (
	"github.com/gin-gonic/gin"

	"gatewayCore/requestHandler"
)

const defaultPort = ":8080"

func main() {

	server := gin.Default()

	// // ===== Train GraphQL Module ===== //

	// server.GET("/train/playground", serviceLinking.PlaygroundHandler("/train"))

	// server.POST("/train", serviceLinking.GraphQLHandler("/train"))

	// // ================================ //

	// // ===== Flight GraphQL Module ==== //

	// server.GET("/flight/playground", serviceLinking.PlaygroundHandler("/flight"))

	// server.POST("/flight", serviceLinking.GraphQLHandler("/flight"))

	// // ================================= //

	// ===== One Gateway Playground ==== //

	// server.GET("/playgroundOld", serviceLinking.PlaygroundHandler("/gateway"))

	// ================================= //

	// ===== One Gateway ======= //

	// server.POST("/gateway", gatewayRouting.RoutingToGQL)

	// ========================= //

	server.Use(requestHandler.Init())

	server.GET("/playground", gin.WrapF(requestHandler.PlaygroundHandler))

	server.POST("/query", gin.WrapF(requestHandler.GraphQLHandler))

	server.POST("/schema", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": requestHandler.MyMergedSchema,
		})
	})

	server.Run(defaultPort)

}
