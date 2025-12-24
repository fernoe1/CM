package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentthree/directmethods"
	"github.com/fernoe1/CM/assignmentthree/iterativemethods"
)

func main() {
	A := [][]float64{
		{4, 1, 2},
		{1, 3, 1},
		{2, 1, 3},
	}

	fmt.Println("Matrix A:")
	for _, row := range A {
		fmt.Println(row)
	}

	B0 := [][]float64{
		{0.0816, 0.0204, 0.0408},
		{0.0204, 0.0612, 0.0204},
		{0.0408, 0.0204, 0.0612},
	}
	tol := 1e-6
	B_iter := iterativemethods.IterativeMethod(A, B0, tol)

	fmt.Println("\nIterative Method - Inverse Approximation:")
	for _, row := range B_iter {
		fmt.Println(row)
	}

	B_lu, err := directmethods.LUFactorization(A)
	if err != nil {
		fmt.Println("LU Method Error:", err)
	} else {
		fmt.Println("\nLU Factorization - Inverse:")
		for _, row := range B_lu {
			fmt.Println(row)
		}
	}

	x0 := []float64{1, 1, 1}
	lambda, vec, err := iterativemethods.PowerMethod(A, x0, tol, 100)
	if err != nil {
		fmt.Println("Power Method Error:", err)
	} else {
		fmt.Println("\nPower Method - Dominant Eigenvalue:", lambda)
		fmt.Println("Corresponding Eigenvector:", vec)
	}

	eigenvalues, eigenvectors := iterativemethods.JacobiMethod(A, tol)
	fmt.Println("\nJacobi Method - Eigenvalues:", eigenvalues)
	fmt.Println("Jacobi Method - Eigenvectors:")
	for _, row := range eigenvectors {
		fmt.Println(row)
	}
}
