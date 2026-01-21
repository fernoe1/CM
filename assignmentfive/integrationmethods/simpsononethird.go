package integrationmethods

import "github.com/fernoe1/CM/assignmentfive/utils"

func SimpsonOneThird(expr string, a, b float64, n int) (float64, uint64) {
	iter := uint64(0)
	f := utils.Function{Expression: expr}
	if n%2 != 0 {
		panic("n must be even for Simpson's 1/3 rule")
	}

	h := (b - a) / float64(n)
	sum := f.F(a) + f.F(b)

	for i := 1; i < n; i++ {
		iter++
		x := a + float64(i)*h
		if i%2 == 0 {
			sum += 2 * f.F(x)
		} else {
			sum += 4 * f.F(x)
		}
	}

	return h * sum / 3, iter
}
