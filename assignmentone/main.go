package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentone/iterativemethods"
)

func main() {
	fmt.Println(iterativemethods.Bisection("x**3 - x - 2", 1, 2, 0.001, 11))
}
