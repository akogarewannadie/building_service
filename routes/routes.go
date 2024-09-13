package routes

import (
	"building_service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/buildings", controllers.CreateBuilding)
	r.GET("/buildings", controllers.GetBuildings)

	return r
}
