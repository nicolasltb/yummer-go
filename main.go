
package main

import (
	"github.com/gin-gonic/gin"
	"yummer-go/database"
	"yummer-go/routes"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	database.InitDB()
	database.SeedData(database.DB)

	routes.SetupRoutes(router)

	router.Run(":8080")
}
