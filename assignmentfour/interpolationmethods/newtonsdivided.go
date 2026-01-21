package interpolationmethods

func NewtonsDivided(x, y []float64, xStar float64) (float64, uint64) {
	iter := uint64(0)
	n := len(x)

	dd := make([][]float64, n)
	for i := range dd {
		dd[i] = make([]float64, n)
		dd[i][0] = y[i]
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			dd[i][j] = (dd[i+1][j-1] - dd[i][j-1]) / (x[i+j] - x[i])
		}
	}

	result := dd[0][0]
	product := 1.0

	for i := 1; i < n; i++ {
		iter++
		product *= (xStar - x[i-1])
		result += dd[0][i] * product
	}

	return result, iter
}
