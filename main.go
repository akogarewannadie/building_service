package main

import (
	"building_service/config"
	"building_service/controllers"
	"building_service/db"
	_ "building_service/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Swagger files handler
	"github.com/swaggo/gin-swagger"        // Swagger middleware
	"time"
)

// @title Building Service API
// @version 1.0
// @description API for managing buildings
// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.LoadConfig()

	time.Sleep(5 * time.Second)

	db.InitDB(cfg)

	r := gin.Default()

	r.POST("/buildings", controllers.CreateBuilding)
	r.GET("/buildings", controllers.GetBuildings)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
