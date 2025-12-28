package iterativemethods

import "github.com/fernoe1/CM/assignmentthree/utils"

func IterativeMethod(A, B [][]float64, tol float64) ([][]float64, int) {
	counter := 0
	I := utils.Identity(len(A))
	E := utils.MatSub(utils.MatMul(A, B), I)

	for utils.MatNorm(E) > tol {
		counter++
		E2 := utils.MatMul(E, E)
		series := utils.MatSub(utils.MatAdd(I, utils.MatMulScalar(E, -1)), E2)
		B = utils.MatMul(B, series)
		E = utils.MatSub(utils.MatMul(A, B), I)
	}

	return B, counter
}
