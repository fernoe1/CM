package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmenttwo/utils"
)

func main() {
	m1 := [][]float64{
		{3.0, 8.0},
		{4.0, 6.0},
	}
	m2 := [][]float64{
		{6.0, 1.0, 1.0},
		{4.0, -2.0, 5.0},
		{2.0, 8.0, 7.0},
	}
	m3 := [][]float64{
		{1.0, 3.0, 1.0, 2.0},
		{5.0, 8.0, 5.0, 3.0},
		{0.0, 4.0, 0.0, 0.0},
		{2.0, 3.0, 2.0, 8.0},
	}

	Matrix1 := utils.Matrix{
		Matrix: m1,
	}

	Matrix2 := utils.Matrix{
		Matrix: m2,
	}

	Matrix3 := utils.Matrix{
		Matrix: m3,
	}

	d1, err := Matrix1.Det()
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("The determinant of", m1, "is", d1)
	}
	d2, err := Matrix2.Det()
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("The determinant of", m2, "is", d2)
	}
	d3, err := Matrix3.Det()
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("The determinant of", m3, "is", d3)
	}
}
