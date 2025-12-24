package directmethods

import "errors"

type LUResult struct {
	L [][]float64
	U [][]float64
}

func luFactorization(A [][]float64) (LUResult, error) {
	n := len(A)
	if n == 0 || len(A[0]) != n {
		return LUResult{}, errors.New("matrix must be square")
	}

	L := make([][]float64, n)
	U := make([][]float64, n)
	for i := 0; i < n; i++ {
		L[i] = make([]float64, n)
		U[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0.0
			for k := 0; k < i; k++ {
				sum += L[i][k] * U[k][j]
			}
			U[i][j] = A[i][j] - sum
		}

		L[i][i] = 1
		for j := i + 1; j < n; j++ {
			sum := 0.0
			for k := 0; k < i; k++ {
				sum += L[j][k] * U[k][i]
			}
			if U[i][i] == 0 {
				return LUResult{}, errors.New("matrix is singular")
			}
			L[j][i] = (A[j][i] - sum) / U[i][i]
		}
	}

	return LUResult{L: L, U: U}, nil
}

func forwardSubstitution(L, B [][]float64) [][]float64 {
	n := len(L)
	X := make([][]float64, n)
	for i := 0; i < n; i++ {
		X[i] = make([]float64, n)
	}

	for col := 0; col < n; col++ {
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < i; j++ {
				sum += L[i][j] * X[j][col]
			}
			X[i][col] = (B[i][col] - sum)
		}
	}

	return X
}

func backwardSubstitution(U, B [][]float64) ([][]float64, error) {
	n := len(U)
	X := make([][]float64, n)
	for i := 0; i < n; i++ {
		X[i] = make([]float64, n)
	}

	for col := 0; col < n; col++ {
		for i := n - 1; i >= 0; i-- {
			sum := 0.0
			for j := i + 1; j < n; j++ {
				sum += U[i][j] * X[j][col]
			}
			if U[i][i] == 0 {
				return nil, errors.New("matrix is singular")
			}
			X[i][col] = (B[i][col] - sum) / U[i][i]
		}
	}

	return X, nil
}

func LUFactorization(A [][]float64) ([][]float64, error) {
	LU, err := luFactorization(A)
	if err != nil {
		return nil, err
	}

	n := len(A)
	I := make([][]float64, n)
	for i := 0; i < n; i++ {
		I[i] = make([]float64, n)
		I[i][i] = 1
	}

	Linv := forwardSubstitution(LU.L, I)

	Uinv, err := backwardSubstitution(LU.U, Linv)
	if err != nil {
		return nil, err
	}

	return Uinv, nil
}
