package interpolationmethods

func NewtonsBackward(x, y []float64, xStar float64) (float64, uint64) {
	iter := uint64(0)
	n := len(x)
	h := x[1] - x[0]

	diff := make([][]float64, n)
	for i := range diff {
		diff[i] = make([]float64, n)
		diff[i][0] = y[i]
	}

	for j := 1; j < n; j++ {
		for i := n - 1; i >= j; i-- {
			diff[i][j] = diff[i][j-1] - diff[i-1][j-1]
		}
	}

	v := (xStar - x[n-1]) / h
	result := diff[n-1][0]
	vTerm := 1.0
	fact := 1.0

	for i := 1; i < n; i++ {
		iter++
		vTerm *= (v + float64(i-1))
		fact *= float64(i)
		result += (vTerm * diff[n-1][i]) / fact
	}

	return result, iter
}
