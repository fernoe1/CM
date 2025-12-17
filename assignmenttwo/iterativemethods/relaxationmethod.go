package iterativemethods

import (
	"errors"
	"math"
)

func RelaxationMethod(a [][]float64, b []float64, omega float64, tol float64, maxIter int) ([]float64, error, int) {
	n := len(a)
	counter := 0

	if n == 0 || len(b) != n {
		return nil, errors.New("invalid matrix dimensions"), counter
	}

	if omega <= 0 || omega >= 2 {
		return nil, errors.New("relaxation factor omega must be in (0, 2)"), counter
	}

	for i := 0; i < n; i++ {
		if a[i][i] == 0 {
			return nil, errors.New("zero found on diagonal, relaxation method not applicable"), counter
		}
	}

	x := make([]float64, n)
	xOld := make([]float64, n)

	for k := 0; k < maxIter; k++ {
		copy(xOld, x)
		counter++

		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += a[i][j] * xOld[j]
				}
			}

			x[i] = (1-omega)*xOld[i] + omega*(b[i]-sum)/a[i][i]
		}

		maxDiff := 0.0
		for i := 0; i < n; i++ {
			diff := math.Abs(x[i] - xOld[i])
			if diff > maxDiff {
				maxDiff = diff
			}
		}

		if maxDiff <= tol {
			return x, nil, counter
		}
	}

	return x, errors.New("relaxation method did not converge within the maximum number of iterations"), counter
}
