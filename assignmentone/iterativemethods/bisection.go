package iterativemethods

import (
	"errors"
	"fmt"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

func Bisection(expr string, a, b, tol float64, maxN int) (float64, error) {
	function := utils.Function{Expression: expr}

	fA := function.F(a)
	fB := function.F(b)

	if fA*fB >= 0 {
		return 0, errors.New("no root in given range")
	}

	n := 0

	for n < maxN {
		c := (a + b) / 2
		fC := function.F(c)

		if math.Abs(fC) < tol || (b-a)/2 < tol {
			return c, nil
		}

		if fA*fC < 0 {
			b = c
			fB = fC
		} else {
			a = c
			fA = fC
		}

		n++
	}

	return 0, fmt.Errorf("no root found after %d iterations", maxN)
}
