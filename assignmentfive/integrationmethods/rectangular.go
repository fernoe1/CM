package integrationmethods

import "github.com/fernoe1/CM/assignmentfive/utils"

func Rectangular(expr string, a, b float64, n int) (float64, uint64) {
	iter := uint64(0)
	f := utils.Function{Expression: expr}
	h := (b - a) / float64(n)

	s := 0.0
	for i := 0; i < n; i++ {
		iter++
		s += f.F(a+float64(i)*h) * h
	}

	return s, iter
}
