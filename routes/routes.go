
package routes

import (
	"github.com/gin-gonic/gin"
	"yummer-go/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/clientes", controllers.GetClientesHTML)
	router.GET("/clientes/new", controllers.CreateClienteHTML)
	router.GET("/clientes/edit/:id", controllers.EditClienteHTML)
	router.POST("/clientes/save", controllers.SaveClienteHTML)
	router.POST("/clientes/delete/:id", controllers.DeleteClienteHTML)

	router.GET("/restaurantes", controllers.GetRestaurantesHTML)
	router.GET("/restaurantes/new", controllers.CreateRestauranteHTML)
	router.GET("/restaurantes/edit/:id", controllers.EditRestauranteHTML)
	router.POST("/restaurantes/save", controllers.SaveRestauranteHTML)
	router.POST("/restaurantes/delete/:id", controllers.DeleteRestauranteHTML)

	router.GET("/mesas", controllers.GetMesasHTML)
	router.GET("/mesas/new", controllers.CreateMesaHTML)
	router.GET("/mesas/edit/:id", controllers.EditMesaHTML)
	router.POST("/mesas/save", controllers.SaveMesaHTML)
	router.POST("/mesas/delete/:id", controllers.DeleteMesaHTML)

	router.GET("/reservas", controllers.GetReservasHTML)
	router.GET("/reservas/new", controllers.CreateReservaHTML)
	router.GET("/reservas/edit/:id", controllers.EditReservaHTML)
	router.POST("/reservas/save", controllers.SaveReservaHTML)
	router.POST("/reservas/delete/:id", controllers.DeleteReservaHTML)
}
