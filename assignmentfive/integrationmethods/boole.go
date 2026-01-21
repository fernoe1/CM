package integrationmethods

import "github.com/fernoe1/CM/assignmentfive/utils"

func Boole(expr string, a, b float64, n int) (float64, uint64) {
	iter := uint64(0)
	f := utils.Function{Expression: expr}
	if n%4 != 0 {
		panic("n must be multiple of 4 for Boole's rule")
	}

	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i <= n; i++ {
		iter++
		x := a + float64(i)*h
		switch {
		case i == 0 || i == n:
			sum += 7 * f.F(x)
		case i%4 == 0:
			sum += 14 * f.F(x)
		case i%2 == 0:
			sum += 12 * f.F(x)
		default:
			sum += 32 * f.F(x)
		}
	}

	return 2 * h * sum / 45, iter
}
