package utils

import "math"

func MatVecMul(A [][]float64, x []float64) []float64 {
	n := len(A)
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < len(x); j++ {
			sum += A[i][j] * x[j]
		}
		result[i] = sum
	}

	return result
}

func MatMul(a, b [][]float64) [][]float64 {
	n := len(a)
	m := len(b[0])
	p := len(b)
	res := make([][]float64, n)

	for i := 0; i < n; i++ {
		res[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			for k := 0; k < p; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res
}

func MatMulScalar(a [][]float64, scalar float64) [][]float64 {
	n := len(a)
	m := len(a[0])
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		res[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			res[i][j] = a[i][j] * scalar
		}
	}

	return res
}

func MatSub(a, b [][]float64) [][]float64 {
	n := len(a)
	m := len(a[0])
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		res[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			res[i][j] = a[i][j] - b[i][j]
		}
	}

	return res
}

func MatAdd(a, b [][]float64) [][]float64 {
	n := len(a)
	m := len(a[0])
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		res[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			res[i][j] = a[i][j] + b[i][j]
		}
	}

	return res
}

func MatNorm(a [][]float64) float64 {
	norm := 0.0
	for i := range a {
		for j := range a[0] {
			if math.Abs(a[i][j]) > norm {
				norm = math.Abs(a[i][j])
			}
		}
	}

	return norm
}

func Identity(n int) [][]float64 {
	I := make([][]float64, n)
	for i := 0; i < n; i++ {
		I[i] = make([]float64, n)
		I[i][i] = 1
	}

	return I
}
