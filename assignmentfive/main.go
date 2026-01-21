package main

import (
	"fmt"
	"math"

	"github.com/fernoe1/CM/assignmentfive/integrationmethods"
)

func main() {
	a := 0.0
	b := math.Pi
	n := 12
	expr := "sin(x)"

	fmt.Println("Integrating f(x) = sin(x) on [0, π]")
	fmt.Println("Exact value = 2.0\n")

	newtonCotesResult, newtonCotesIter := integrationmethods.NewtonCotes(expr, a, b, n)
	trapezoidalResult, trapezoidalIter := integrationmethods.Trapezoidal(expr, a, b, n)
	simpsonOneThirdResult, simpsonOneThirdIter := integrationmethods.SimpsonOneThird(expr, a, b, n)
	simpsonThreeEighthResult, simpsonThreeEighthIter := integrationmethods.SimpsonThreeEight(expr, a, b, n)
	booleResult, booleIter := integrationmethods.Boole(expr, a, b, n)
	weddleResult, weddleIter := integrationmethods.Weddle(expr, a, b, n)

	fmt.Println("Newton–Cotes:", newtonCotesResult, "Iterations:", newtonCotesIter)
	fmt.Println("Trapezoidal:", trapezoidalResult, "Iterations:", trapezoidalIter)
	fmt.Println("Simpson 1/3:", simpsonOneThirdResult, "Iterations:", simpsonOneThirdIter)
	fmt.Println("Simpson 3/8:", simpsonThreeEighthResult, "Iterations:", simpsonThreeEighthIter)
	fmt.Println("Boole:", booleResult, "Iterations:", booleIter)
	fmt.Println("Weddle:", weddleResult, "Iterations:", weddleIter)
}
