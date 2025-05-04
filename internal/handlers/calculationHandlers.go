package handlers

import (
	myResp "back/http"
	"back/internal/calculationService"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, myResp.Error(fmt.Sprintf("%v", err)))
	}
	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculations(c echo.Context) error {
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	calculation, err := h.service.CreateCalculation(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}
	return c.JSON(http.StatusCreated, calculation)
}

func (h *CalculationHandler) PatchCalculations(c echo.Context) error {
	id := c.Param("id")

	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}

	calculation, err := h.service.UpdateCalculation(id, req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}
	return c.JSON(http.StatusOK, calculation)
}

func (h *CalculationHandler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")

	err := h.service.DeleteCalculation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, myResp.Error(fmt.Sprintf("%v", err)))
	}
	return c.NoContent(http.StatusNoContent)
}
