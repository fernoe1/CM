package iterativemethods

import (
	"errors"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

const maxIterations = 1000

func FixedPoint(expr string, x0, tol float64) (float64, error, int) {
	iterations := 0
	function := utils.Function{Expression: expr}

	for i := 0; i < maxIterations; i++ {
		x1 := function.F(x0)

		if math.Abs(x1-x0) <= tol {
			return x1, nil, iterations
		}

		x0 = x1
		iterations++
	}

	return 0, errors.New("could not converge to root after 1000 iterations"), iterations
}
