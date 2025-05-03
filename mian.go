package main

import (
	"back/calculate"
	"back/database"
	myResp "back/http"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, getAllCalculations())
}

func postCalculations(c echo.Context) error {

	var req calculate.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	result, err := calculate.CalculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	calc := calculate.Calculations{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     result,
	}

	dbAnswer := database.DB.Create(&calc)

	if dbAnswer.Error != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", dbAnswer.Error)))
	}

	return c.JSON(http.StatusCreated, calc)
}

func deleteCalculations(c echo.Context) error {
	id := c.Param("id")

	var calc calculate.Calculations

	dbResult := database.DB.Where("id = ?", id).First(&calc)
	if dbResult.Error != nil {
		return c.JSON(http.StatusNotFound, myResp.Error(fmt.Sprintf("%v", dbResult.Error)))
	}

	database.DB.Delete(&calc)

	return c.NoContent(http.StatusNoContent)
}

func patchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req calculate.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	result, err := calculate.CalculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	var calc calculate.Calculations

	fmt.Printf("Извлеченный ID: %s\n", id)

	dbResult := database.DB.Where("id = ?", id).First(&calc)
	if dbResult.Error != nil {
		return c.JSON(http.StatusNotFound, myResp.Error(fmt.Sprintf("%v", dbResult.Error)))
	}

	fmt.Printf("Извлеченный ID 2 : %s\n", id)

	calc.Expression = req.Expression
	calc.Result = result

	database.DB.Save(&calc)

	return c.JSON(http.StatusOK, getAllCalculations())
}

func getAllCalculations() []calculate.Calculations {
	var calculations []calculate.Calculations

	database.DB.Find(&calculations)

	return calculations
}

func main() {

	database.InitDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)
	e.DELETE("/calculations/:id", deleteCalculations)
	e.PATCH("/calculations/:id", patchCalculations)

	e.Start("localhost:8080")
}
