package controllers

import (
	"building_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBuilding godoc
// @Summary Create a new building
// @Description Add a new building to the database
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body models.Building true "Building"
// @Success 200 {object} CreateBuildingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /buildings [post]
func CreateBuilding(c *gin.Context) {
	var building models.Building

	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
		return
	}

	// Вставка данных в базу через модель
	err := models.CreateBuilding(building)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create building"})
		return
	}

	c.JSON(http.StatusOK, CreateBuildingResponse{Message: "Building created successfully"})
}

// GetBuildings godoc
// @Summary Get list of buildings
// @Description Get list of buildings with optional filters
// @Tags buildings
// @Accept json
// @Produce json
// @Param city query string false "City"
// @Param year query int false "Year"
// @Param floors query int false "Number of floors"
// @Success 200 {object} GetBuildingsResponse
// @Failure 500 {object} ErrorResponse
// @Router /buildings [get]
func GetBuildings(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year")
	floors := c.Query("floors")

	buildings, err := models.GetBuildings(city, year, floors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to retrieve buildings"})
		return
	}

	c.JSON(http.StatusOK, GetBuildingsResponse{Buildings: buildings})
}

type CreateBuildingResponse struct {
	Message string `json:"message"`
}

type GetBuildingsResponse struct {
	Buildings []models.Building `json:"buildings"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
