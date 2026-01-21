package interpolationmethods

func Spline(x, y []float64, xStar float64) (float64, uint64) {
	iter := uint64(0)
	n := len(x)
	h := make([]float64, n-1)

	for i := 0; i < n-1; i++ {
		h[i] = x[i+1] - x[i]
	}

	alpha := make([]float64, n)
	for i := 1; i < n-1; i++ {
		alpha[i] = (3/h[i])*(y[i+1]-y[i]) - (3/h[i-1])*(y[i]-y[i-1])
	}

	l := make([]float64, n)
	mu := make([]float64, n)
	z := make([]float64, n)

	l[0] = 1
	for i := 1; i < n-1; i++ {
		l[i] = 2*(x[i+1]-x[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}
	l[n-1] = 1

	c := make([]float64, n)
	b := make([]float64, n-1)
	d := make([]float64, n-1)

	for j := n - 2; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
		b[j] = (y[j+1]-y[j])/h[j] - h[j]*(c[j+1]+2*c[j])/3
		d[j] = (c[j+1] - c[j]) / (3 * h[j])
	}

	i := 0
	for i < n-1 && xStar > x[i+1] {
		iter++
		i++
	}

	dx := xStar - x[i]
	return y[i] + b[i]*dx + c[i]*dx*dx + d[i]*dx*dx*dx, iter
}
