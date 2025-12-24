package iterativemethods

import (
	"math"

	"github.com/fernoe1/CM/assignmentthree/utils"
)

func maxOffDiagonal(A [][]float64) (int, int, float64) {
	n := len(A)
	maxVal := 0.0
	p, q := 0, 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if math.Abs(A[i][j]) > maxVal {
				maxVal = math.Abs(A[i][j])
				p, q = i, j
			}
		}
	}

	return p, q, maxVal
}

func JacobiMethod(A [][]float64, tol float64) (eigenvalues []float64, eigenvectors [][]float64) {
	n := len(A)
	V := utils.Identity(n)

	D := make([][]float64, n)
	for i := 0; i < n; i++ {
		D[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			D[i][j] = A[i][j]
		}
	}

	for {
		p, q, maxOff := maxOffDiagonal(D)
		if maxOff < tol {
			break
		}

		if D[p][p] == D[q][q] {
			theta := math.Pi / 4
			c := math.Cos(theta)
			s := math.Sin(theta)
			applyRotation(D, V, p, q, c, s)
		} else {
			tau := (2 * D[p][q]) / (D[q][q] - D[p][p])
			t := math.Copysign(1.0/(math.Abs(tau)+math.Sqrt(1+tau*tau)), tau)
			c := 1 / math.Sqrt(1+t*t)
			s := t * c
			applyRotation(D, V, p, q, c, s)
		}
	}

	eigenvalues = make([]float64, n)
	for i := 0; i < n; i++ {
		eigenvalues[i] = D[i][i]
	}

	return eigenvalues, V
}

func applyRotation(A, V [][]float64, p, q int, c, s float64) {
	n := len(A)
	for i := 0; i < n; i++ {
		if i != p && i != q {
			Aip := A[i][p]
			Aiq := A[i][q]
			A[i][p] = c*Aip - s*Aiq
			A[p][i] = A[i][p]
			A[i][q] = c*Aiq + s*Aip
			A[q][i] = A[i][q]
		}
	}

	App := A[p][p]
	Aqq := A[q][q]
	Apq := A[p][q]

	A[p][p] = c*c*App - 2*s*c*Apq + s*s*Aqq
	A[q][q] = s*s*App + 2*s*c*Apq + c*c*Aqq
	A[p][q] = 0
	A[q][p] = 0

	for i := 0; i < n; i++ {
		Vip := V[i][p]
		Viq := V[i][q]
		V[i][p] = c*Vip - s*Viq
		V[i][q] = s*Vip + c*Viq
	}
}
