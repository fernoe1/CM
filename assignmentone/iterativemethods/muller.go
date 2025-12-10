package iterativemethods

import (
	"fmt"
	"math"

	"github.com/fernoe1/CM/assignmentone/utils"
)

func Muller(expr string, x0, x1, x2, tol float64, maxN int) (float64, error) {
	function := utils.Function{Expression: expr}
	n := 0

	for n < maxN {
		y0 := function.F(x0)
		y1 := function.F(x1)
		y2 := function.F(x2)

		h0 := x1 - x0
		h1 := x2 - x1
		d0 := (y1 - y0) / h0
		d1 := (y2 - y1) / h1
		a := (d1 - d0) / (h1 + h0)
		b := a*h1 + d1
		c := y2

		d := math.Sqrt(b*b - 4*a*c)

		var denom float64
		if math.Abs(b+d) > math.Abs(b-d) {
			denom = b + d
		} else {
			denom = b - d
		}

		xn := x2 - (2*c)/denom

		if math.Abs(function.F(xn)) < tol {
			return xn, nil
		}

		x0, x1, x2 = x1, x2, xn
		n++
	}

	return 0, fmt.Errorf("no root found after %d iterations", maxN)
}
