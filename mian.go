package main

import (
	"back/calculate"
	myResp "back/http"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var calculations = []calculate.CalculationType{}

func getCalculations(c echo.Context) error {
	return c.JSON(http.StatusOK, calculations)
}

func postCalculations(c echo.Context) error {

	var req calculate.CalculationRequestType

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	result, err := calculate.CalculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	calc := saveToDB(req.Expression, result)

	return c.JSON(http.StatusCreated, calc)
}

func deleteCalculations(c echo.Context) error {
	id := c.Param("id")

	for i, v := range calculations {
		if v.ID == id {
			calculations = append(calculations[:i], calculations[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, myResp.Error("Calculation not found"))
}

func patchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req calculate.CalculationRequestType

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	result, err := calculate.CalculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	for i, v := range calculations {
		if v.ID == id {
			calculations[i].Expression = req.Expression
			calculations[i].Result = result
			return c.JSON(http.StatusOK, calculations)
		}
	}
	return c.JSON(http.StatusBadRequest, myResp.Error("Calculation not found"))
}

func saveToDB(expressions string, result string) calculate.CalculationType {

	calc := calculate.CalculationType{
		ID:         uuid.NewString(),
		Expression: expressions,
		Result:     result,
	}

	calculations = append(calculations, calc)

	return calc
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)
	e.DELETE("/calculations/:id", deleteCalculations)
	e.PATCH("/calculations/:id", patchCalculations)

	e.Start("localhost:8080")
}
