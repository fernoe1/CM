package integrationmethods

import "github.com/fernoe1/CM/assignmentfive/utils"

func SimpsonThreeEight(expr string, a, b float64, n int) (float64, uint64) {
	iter := uint64(0)
	f := utils.Function{Expression: expr}
	if n%3 != 0 {
		panic("n must be multiple of 3 for Simpson's 3/8 rule")
	}

	h := (b - a) / float64(n)
	sum := f.F(a) + f.F(b)

	for i := 1; i < n; i++ {
		iter++
		x := a + float64(i)*h
		if i%3 == 0 {
			sum += 2 * f.F(x)
		} else {
			sum += 3 * f.F(x)
		}
	}

	return 3 * h * sum / 8, iter
}
