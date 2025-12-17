package directmethods

import (
	"errors"

	"github.com/fernoe1/CM/assignmenttwo/utils"
)

func CramersMethod(a [][]float64, b []float64) ([]float64, error) {
	matrix := utils.Matrix{
		Matrix: a,
	}

	D, err := matrix.Det()
	if err != nil {
		return nil, err
	}

	if D == 0 {
		return nil, errors.New("system is either singular or not compatible")
	}

	n := len(a)
	x := make([]float64, n)

	for i := 0; i < n; i++ {
		Ai := make([][]float64, n)
		for r := 0; r < n; r++ {
			Ai[r] = make([]float64, n)
			copy(Ai[r], a[r])
		}

		for r := 0; r < n; r++ {
			Ai[r][i] = b[r]
		}

		modMatrix := utils.Matrix{
			Matrix: Ai,
		}

		Di, err := modMatrix.Det()
		if err != nil {
			return nil, err
		}

		x[i] = Di / D
	}

	return x, nil
}
