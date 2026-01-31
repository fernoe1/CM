package numericalmethods

import "github.com/fernoe1/CM/assignmentsix/utils"

func TaylorSeries(expr string, x0, y0, x float64, order int) float64 {
	f := utils.Function{Expression: expr}
	h := x - x0
	y := y0

	factorial := 1.0
	power := h

	for k := 1; k <= order; k++ {
		if k > 1 {
			factorial *= float64(k)
			power *= h
		}

		derivative := f.NthDerivative(x0, k-1)
		y += power * derivative / factorial
	}

	return y
}
