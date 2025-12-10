package utils

import "github.com/Knetic/govaluate"

type Function struct {
	Expression string
}

func (f Function) F(x float64) float64 {
	expression, _ := govaluate.NewEvaluableExpression(f.Expression)
	parameters := map[string]interface{}{"x": x}
	result, _ := expression.Evaluate(parameters)

	return result.(float64)
}
