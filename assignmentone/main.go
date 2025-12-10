package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentone/iterativemethods"
)

func main() {
	fmt.Println(iterativemethods.Bisection("x**3 - x - 2", 1, 2, 1e-3, 13))

	fmt.Println(iterativemethods.FixedPoint("cos(x)", 0.5, 1e-6))

	fmt.Println(iterativemethods.NewtonRaphson("x**3 - 2*x - 5", "3*x**2 - 2", 2, 1e-7))

	fmt.Println(iterativemethods.Secant("x**3 - 2*x - 5", 2, 3, 1e-6, 1000))

	fmt.Println(iterativemethods.FalsePosition("2 * exp(x) * sin(x) - 3", 0, 1, 1e-6, 1000))

	fmt.Println(iterativemethods.Muller("cos(x) - x * exp(x)", -1, 0, 1, 1e-6, 10000))
}
