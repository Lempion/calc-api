package calculationService

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculations, error)
	GetAllCalculation() ([]Calculations, error)
	GetCalculationByID(id string) (Calculations, error)
	UpdateCalculation(id string, expression string) (Calculations, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) CreateCalculation(expression string) (Calculations, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculations{}, err
	}

	calc := Calculations{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculations{}, err
	}

	return calc, nil
}

func (s *calcService) GetAllCalculation() ([]Calculations, error) {
	return s.repo.GetAllCalculation()
}

func (s *calcService) GetCalculationByID(id string) (Calculations, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *calcService) UpdateCalculation(id string, expression string) (Calculations, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculations{}, err
	}

	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculations{}, err
	}

	calc.Expression = expression
	calc.Result = result

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculations{}, err
	}

	return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	evaluate, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", evaluate), err
}
