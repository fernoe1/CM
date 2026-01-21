package integrationmethods

import "github.com/fernoe1/CM/assignmentfive/utils"

func Weddle(expr string, a, b float64, n int) (float64, uint64) {
	iter := uint64(0)
	f := utils.Function{Expression: expr}
	if n%6 != 0 {
		panic("n must be multiple of 6 for Weddle's rule")
	}

	h := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i <= n; i++ {
		iter++
		x := a + float64(i)*h
		switch {
		case i == 0 || i == n:
			sum += f.F(x)
		case i%6 == 3:
			sum += 6 * f.F(x)
		case i%2 == 1:
			sum += 5 * f.F(x)
		default:
			sum += f.F(x)
		}
	}

	return 3 * h * sum / 10, iter
}
