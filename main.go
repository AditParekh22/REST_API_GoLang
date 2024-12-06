package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/events", getEvets)
	server.Run(":8080")

}

func getEvets(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "Heloo"})
}
