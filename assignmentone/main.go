package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentone/iterativemethods"
)

func main() {
	fmt.Println(iterativemethods.Bisection("x**3 - x - 2", 1, 2, 1e-3, 11))

	fmt.Println(iterativemethods.FixedPoint("cos(x)", 0.5, 1e-6))
}
