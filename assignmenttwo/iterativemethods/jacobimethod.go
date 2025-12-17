package iterativemethods

import (
	"errors"
	"math"
)

func JacobiMethod(a [][]float64, b []float64, tol float64, maxIter int) ([]float64, error, int) {
	n := len(a)
	counter := 0

	if n == 0 || len(b) != n {
		return nil, errors.New("invalid matrix dimensions"), counter
	}

	for i := 0; i < n; i++ {
		if a[i][i] == 0 {
			return nil, errors.New("zero found on diagonal, jacobi method not applicable"), counter
		}
	}
	xOld := make([]float64, n)
	xNew := make([]float64, n)

	for k := 0; k < maxIter; k++ {
		counter++
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += a[i][j] * xOld[j]
				}
			}
			xNew[i] = (b[i] - sum) / a[i][i]
		}

		maxDiff := 0.0
		for i := 0; i < n; i++ {
			diff := math.Abs(xNew[i] - xOld[i])
			if diff > maxDiff {
				maxDiff = diff
			}
			xOld[i] = xNew[i]
		}

		if maxDiff <= tol {
			return xNew, nil, counter
		}
	}

	return xNew, errors.New("jacobi method did not converge within the maximum number of iterations"), counter
}
