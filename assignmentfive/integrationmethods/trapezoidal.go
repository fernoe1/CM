package integrationmethods

import "github.com/fernoe1/CM/assignmentfive/utils"

func Trapezoidal(expr string, a, b float64, n int) (float64, uint64) {
	iter := uint64(0)
	f := utils.Function{Expression: expr}
	h := (b - a) / float64(n)
	sum := f.F(a) + f.F(b)

	for i := 1; i < n; i++ {
		iter++
		sum += 2 * f.F(a+float64(i)*h)
	}

	return h * sum / 2, iter
}
