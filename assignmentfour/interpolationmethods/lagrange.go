package interpolationmethods

func Lagrange(x []float64, y []float64, xStar float64) (float64, uint64) {
	iter := uint64(0)
	n := len(x)
	result := 0.0

	for i := 0; i < n; i++ {
		iter++
		prod := 1.0

		for j := 0; j < n; j++ {
			if i != j {
				prod *= (xStar - x[j]) / (x[i] - x[j])
			}
		}

		result += prod * y[i]
	}

	return result, iter
}
