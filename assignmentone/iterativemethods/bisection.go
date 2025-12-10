package iterativemethods

import (
	"errors"
	"fmt"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

func Bisection(expr string, a, b, tol float64, maxN int) (float64, error, int) {
	iterations := 0
	function := utils.Function{Expression: expr}
	n := 1
	fA := function.F(a)
	fB := function.F(b)

	if fA*fB >= 0 {
		return 0, errors.New("no root in given range"), iterations
	}

	for n <= maxN {
		c := (a + b) / 2
		fC := function.F(c)

		if math.Abs(fC) < tol || (b-a)/2 < tol {
			return c, nil, iterations
		}

		if fA*fC < 0 {
			b, fB = c, fC
		} else {
			a, fA = c, fC
		}

		n++
		iterations++
	}

	return 0, fmt.Errorf("no root found after %d iterations", maxN), iterations
}
