package directmethods

import (
	"errors"
	"math"
)

func GaussJordanMethod(a [][]float64, b []float64) ([]float64, error) {
	n := len(a)

	if n == 0 || len(b) != n {
		return nil, errors.New("invalid matrix dimensions")
	}

	aug := make([][]float64, n)
	for i := 0; i < n; i++ {
		if len(a[i]) != n {
			return nil, errors.New("matrix must be square")
		}
		aug[i] = make([]float64, n+1)
		copy(aug[i], a[i])
		aug[i][n] = b[i]
	}

	for k := 0; k < n; k++ {
		pivotRow := k
		for i := k + 1; i < n; i++ {
			if math.Abs(aug[i][k]) > math.Abs(aug[pivotRow][k]) {
				pivotRow = i
			}
		}

		if aug[pivotRow][k] == 0 {
			return nil, errors.New("system is singular or not compatible")
		}

		if pivotRow != k {
			aug[k], aug[pivotRow] = aug[pivotRow], aug[k]
		}

		pivot := aug[k][k]
		for j := k; j <= n; j++ {
			aug[k][j] /= pivot
		}

		for i := 0; i < n; i++ {
			if i == k {
				continue
			}

			factor := aug[i][k]
			for j := k; j <= n; j++ {
				aug[i][j] -= factor * aug[k][j]
			}
		}
	}

	x := make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = aug[i][n]
	}

	return x, nil
}
