package main

import (
	"back/internal/calculationService"
	"back/internal/database"
	"back/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	calcRepo := calculationService.NewCalculationRepository(db)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}
