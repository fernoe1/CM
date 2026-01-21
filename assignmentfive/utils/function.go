package utils

import (
	"math"

	"github.com/Knetic/govaluate"
)

type Function struct {
	Expression string
}

func (f Function) F(x float64) float64 {
	functions := map[string]govaluate.ExpressionFunction{
		"cos": func(args ...interface{}) (interface{}, error) {
			return math.Cos(args[0].(float64)), nil
		},
		"sin": func(args ...interface{}) (interface{}, error) {
			return math.Sin(args[0].(float64)), nil
		},
		"tan": func(args ...interface{}) (interface{}, error) {
			return math.Tan(args[0].(float64)), nil
		},
		"exp": func(args ...interface{}) (interface{}, error) {
			return math.Exp(args[0].(float64)), nil
		},
		"sqrt": func(args ...interface{}) (interface{}, error) {
			return math.Sqrt(args[0].(float64)), nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(f.Expression, functions)
	if err != nil {
		panic(err)
	}

	parameters := map[string]interface{}{"x": x}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		panic(err)
	}

	return result.(float64)
}

func (f Function) DerivativeF(x float64) float64 {
	h := 1e-5

	return (f.F(x+h) - f.F(x-h)) / (2 * h)
}
