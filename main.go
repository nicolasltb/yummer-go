
package main

import (
	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/routes"
)

func main() {
	router := gin.Default()

	database.InitDB()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
