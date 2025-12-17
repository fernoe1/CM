package directmethods

import (
	"errors"
	"math"
)

func GaussianMethod(a [][]float64, b []float64) ([]float64, error) {
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

	for i := 0; i < n; i++ {
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(aug[k][i]) > math.Abs(aug[maxRow][i]) {
				maxRow = k
			}
		}

		if aug[maxRow][i] == 0 {
			return nil, errors.New("system is singular or not compatible")
		}

		aug[i], aug[maxRow] = aug[maxRow], aug[i]

		for k := i + 1; k < n; k++ {
			factor := aug[k][i] / aug[i][i]
			for j := i; j <= n; j++ {
				aug[k][j] -= factor * aug[i][j]
			}
		}
	}

	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := aug[i][n]
		for j := i + 1; j < n; j++ {
			sum -= aug[i][j] * x[j]
		}
		x[i] = sum / aug[i][i]
	}

	return x, nil
}
