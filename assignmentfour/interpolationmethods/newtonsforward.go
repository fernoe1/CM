package interpolationmethods

func NewtonsForward(x, y []float64, xStar float64) (float64, uint64) {
	iter := uint64(0)
	n := len(x)
	h := x[1] - x[0]

	diff := make([][]float64, n)
	for i := range diff {
		diff[i] = make([]float64, n)
		diff[i][0] = y[i]
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			diff[i][j] = diff[i+1][j-1] - diff[i][j-1]
		}
	}

	u := (xStar - x[0]) / h
	result := diff[0][0]
	uTerm := 1.0
	fact := 1.0

	for i := 1; i < n; i++ {
		iter++
		uTerm *= (u - float64(i-1))
		fact *= float64(i)
		result += (uTerm * diff[0][i]) / fact
	}

	return result, iter
}
