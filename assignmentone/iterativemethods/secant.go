package iterativemethods

import (
	"fmt"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

func Secant(expr string, x0, x1, tol float64, maxN int) (float64, error, int) {
	iterations := 0
	function := utils.Function{Expression: expr}
	n := 1
	f0 := function.F(x0)
	f1 := function.F(x1)

	for n <= maxN {
		xn := x1 - f1*((x1-x0)/(f1-f0))
		fn := function.F(xn)

		if math.Abs(fn) < tol {
			return xn, nil, iterations
		}

		x0, f0, x1, f1 = x1, f1, xn, fn

		n++
		iterations++
	}

	return 0, fmt.Errorf("no root found after %d iterations", maxN), iterations
}
