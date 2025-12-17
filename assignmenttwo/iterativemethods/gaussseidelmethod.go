package iterativemethods

import (
	"errors"
	"math"
)

func GaussSeidelMethod(a [][]float64, b []float64, tol float64, maxIter int) ([]float64, error, int) {
	n := len(a)
	counter := 0

	if n == 0 || len(b) != n {
		return nil, errors.New("invalid matrix dimensions"), counter
	}

	for i := 0; i < n; i++ {
		if a[i][i] == 0 {
			return nil, errors.New("zero found on diagonal, Gauss-Seidel not applicable"), counter
		}
	}

	x := make([]float64, n)
	xOld := make([]float64, n)

	for k := 0; k < maxIter; k++ {
		counter++
		copy(xOld, x)

		for i := 0; i < n; i++ {
			sum1 := 0.0
			for j := 0; j < i; j++ {
				sum1 += a[i][j] * x[j]
			}

			sum2 := 0.0
			for j := i + 1; j < n; j++ {
				sum2 += a[i][j] * xOld[j]
			}

			x[i] = (b[i] - sum1 - sum2) / a[i][i]
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

	return x, errors.New("gauss-seidel method did not converge within the maximum number of iterations"), counter
}
