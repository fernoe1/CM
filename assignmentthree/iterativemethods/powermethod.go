package iterativemethods

import (
	"errors"
	"math"

	"github.com/fernoe1/CM/assignmentthree/utils"
)

func normalize(x []float64) []float64 {
	max := 0.0
	for _, v := range x {
		if math.Abs(v) > max {
			max = math.Abs(v)
		}
	}
	if max == 0 {

		return x
	}
	normed := make([]float64, len(x))
	for i, v := range x {
		normed[i] = v / max
	}

	return normed
}

func rayleighQuotient(A [][]float64, x []float64) float64 {
	n := len(x)
	num := 0.0
	den := 0.0
	for i := 0; i < n; i++ {
		temp := 0.0
		for j := 0; j < n; j++ {
			temp += A[i][j] * x[j]
		}
		num += x[i] * temp
		den += x[i] * x[i]
	}
	if den == 0 {

		return 0
	}

	return num / den
}

func PowerMethod(A [][]float64, x []float64, tol float64, maxN uint64) (float64, []float64, error) {
	n := len(A)
	if n == 0 || len(A[0]) != n {
		return 0, nil, errors.New("matrix must be square")
	}
	if len(x) != n {
		return 0, nil, errors.New("initial vector size mismatch")
	}

	xk := normalize(x)
	var lambda float64

	for k := uint64(0); k < maxN; k++ {
		y := utils.MatVecMul(A, xk)

		xk1 := normalize(y)

		lambdaNext := rayleighQuotient(A, xk1)

		diff := 0.0
		for i := 0; i < n; i++ {
			d := xk1[i] - xk[i]
			if math.Abs(d) > diff {
				diff = math.Abs(d)
			}
		}
		if diff < tol {
			return lambdaNext, xk1, nil
		}

		xk = xk1
		lambda = lambdaNext
	}

	return lambda, xk, nil
}
