package calculate

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func CalculateExpression(expression string) (string, error) {
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
