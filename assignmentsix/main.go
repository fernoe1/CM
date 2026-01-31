package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentsix/numericalmethods"
)

func main() {
	y1 := numericalmethods.Euler("x + exp(x)", 0, 1, 0.2, 10)
	y2 := numericalmethods.ModifiedEuler("x + exp(x)", 0, 1, 0.2, 10)
	y3 := numericalmethods.RK3("x + exp(x)", 0, 1, 0.2, 10)
	y4 := numericalmethods.RK4("x + exp(x)", 0, 1, 0.2, 10)
	y5 := numericalmethods.TaylorSeries("x + exp(x)", 0, 1, 0.2, 4)

	fmt.Println("Euler:", y1)
	fmt.Println("ModifiedEuler:", y2)
	fmt.Println("RK3:", y3)
	fmt.Println("RK4:", y4)
	fmt.Println("TaylorSeries:", y5)
}
