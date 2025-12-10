package iterativemethods

import (
	"errors"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

const maxN = 1000

func NewtonRaphson(expr, dExpr string, x0, tol float64) (float64, error) {
	function := utils.Function{Expression: expr}
	dFunction := utils.Function{Expression: dExpr}

	for i := 0; i < maxN; i++ {
		x1 := x0 - (function.F(x0) / dFunction.F(x0))

		if math.Abs(x1-x0) <= tol {
			return x1, nil
		}

		x0 = x1
	}

	return 0, errors.New("could not converge to root after 1000 iterations")
}
