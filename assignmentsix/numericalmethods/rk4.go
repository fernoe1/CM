package numericalmethods

import "github.com/fernoe1/CM/assignmentsix/utils"

func RK4(expr string, x0, y0, x float64, steps int) float64 {
	f := utils.Function{Expression: expr}

	h := (x - x0) / float64(steps)
	y := y0
	xn := x0

	for i := 0; i < steps; i++ {
		k1 := f.F(xn)
		k2 := f.F(xn + h/2)
		k3 := f.F(xn + h/2)
		k4 := f.F(xn + h)

		y += (h / 6) * (k1 + 2*k2 + 2*k3 + k4)
		xn += h
	}

	return y
}
