package routes

import (
	"github.com/argalon/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes for the service
func SetupRouter() *gin.Engine {
	r := gin.Default()

	order := r.Group("/argalon/")
	{
		order.POST("divide_two", controllers.DivideTwoController)

	}

	return r
}
