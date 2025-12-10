package iterativemethods

import (
	"fmt"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

func FalsePosition(expr string, a, b, tol float64, maxN int) (float64, error, int) {
	iterations := 0
	function := utils.Function{Expression: expr}
	n := 0
	fA := function.F(a)
	fB := function.F(b)

	for n < maxN {
		c := (a*fB - b*fA) / (fB - fA)
		fC := function.F(c)

		if math.Abs(fC) < tol {
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
