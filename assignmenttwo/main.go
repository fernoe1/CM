package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmenttwo/directmethods"
	"github.com/fernoe1/CM/assignmenttwo/iterativemethods"
)

func main() {
	A := [][]float64{
		{10, 2, 1},
		{2, 10, 3},
		{1, 3, 10},
	}
	b := []float64{9, 15, 14}
	tol := 1e-6
	maxIter := 1000
	omega := 1.0

	xCramer, err := directmethods.CramersMethod(A, b)
	if err != nil {
		fmt.Println("Cramer Error:", err)
	} else {
		fmt.Printf("Cramer: %v\n", xCramer)
	}

	xGauss, err := directmethods.GaussianMethod(A, b)
	if err != nil {
		fmt.Println("Gaussian Error:", err)
	} else {
		fmt.Printf("Gaussian: %v\n", xGauss)
	}

	xGaussJordan, err := directmethods.GaussJordanMethod(A, b)
	if err != nil {
		fmt.Println("Gauss-Jordan Error:", err)
	} else {
		fmt.Printf("Gauss-Jordan: %v\n", xGaussJordan)
	}

	xJacobi, err, iterJacobi := iterativemethods.JacobiMethod(A, b, tol, maxIter)
	if err != nil {
		fmt.Println("Jacobi Error:", err)
	} else {
		fmt.Printf("Jacobi: %v\n", xJacobi)
		fmt.Print("Iterations: ", iterJacobi, "\n")
	}

	xGS, err, iterGaussSeidel := iterativemethods.GaussSeidelMethod(A, b, tol, maxIter)
	if err != nil {
		fmt.Println("Gauss-Seidel Error:", err)
	} else {
		fmt.Printf("Gauss-Seidel: %v\n", xGS)
		fmt.Print("Iterations: ", iterGaussSeidel, "\n")
	}

	xSOR, err, iterRelaxation := iterativemethods.RelaxationMethod(A, b, omega, tol, maxIter)
	if err != nil {
		fmt.Println("Relaxation Error:", err)
	} else {
		fmt.Printf("Relaxation: %v\n", xSOR)
		fmt.Print("Iterations: ", iterRelaxation, "\n")
	}
}
