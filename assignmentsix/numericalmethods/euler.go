package numericalmethods

import "github.com/fernoe1/CM/assignmentsix/utils"

func Euler(expr string, x0, y0, x float64, steps int) float64 {
	f := utils.Function{Expression: expr}

	h := (x - x0) / float64(steps)
	y := y0
	xn := x0

	for i := 0; i < steps; i++ {
		y += h * f.F(xn)
		xn += h
	}

	return y
}
