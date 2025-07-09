package routes

import (
	"github.com/gin-gonic/gin"
	"yummer-go/controllers"
)

func SetupRoutes(router *gin.Engine) {
	clienteRoutes := router.Group("/clientes")
	{
		clienteRoutes.POST("/", controllers.CreateCliente)
		clienteRoutes.GET("/", controllers.GetClientes)
		clienteRoutes.GET("/id", controllers.GetCliente)
		clienteRoutes.PUT("/id", controllers.UpdateCliente)
		clienteRoutes.DELETE("/id", controllers.DeleteCliente)
	}

	restauranteRoutes := router.Group("/restaurantes")
	{
		restauranteRoutes.POST("/", controllers.CreateRestaurante)
		restauranteRoutes.GET("/", controllers.GetRestaurantes)
		restauranteRoutes.GET("/id", controllers.GetRestaurante)
		restauranteRoutes.PUT("/id", controllers.UpdateRestaurante)
		restauranteRoutes.DELETE("/id", controllers.DeleteRestaurante)
	}

	mesaRoutes := router.Group("/mesas")
	{
		mesaRoutes.POST("/", controllers.CreateMesa)
		mesaRoutes.GET("/", controllers.GetMesas)
		mesaRoutes.GET("/id", controllers.GetMesa)
		mesaRoutes.PUT("/id", controllers.UpdateMesa)
		mesaRoutes.DELETE("/id", controllers.DeleteMesa)
	}

	reservaRoutes := router.Group("/reservas")
	{
		reservaRoutes.POST("/", controllers.CreateReserva)
		reservaRoutes.GET("/", controllers.GetReservas)
		reservaRoutes.GET("/id", controllers.GetReserva)
		reservaRoutes.PUT("/id", controllers.UpdateReserva)
		reservaRoutes.DELETE("/id", controllers.DeleteReserva)
	}
}