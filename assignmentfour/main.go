package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentfour/interpolationmethods"
)

func main() {
	x := []float64{1, 2, 3, 4}
	y := []float64{1, 4, 9, 16}
	xStar := 2.5

	fmt.Println("Interpolating f(x) = x^2 at x =", xStar)
	fmt.Println("Exact value:", 6.25)
	fmt.Println()

	lagrangeX, lagrangeIter := interpolationmethods.Lagrange(x, y, xStar)
	fmt.Println("Lagrange:", lagrangeX, "Iterations:", lagrangeIter)
	newtonfX, newtonfIter := interpolationmethods.NewtonsForward(x, y, xStar)
	fmt.Println("Newton's forward:", newtonfX, "Iterations:", newtonfIter)
	newtonbX, newtonbIter := interpolationmethods.NewtonsBackward(x, y, xStar)
	fmt.Println("Newton's backward:", newtonbX, "Iterations:", newtonbIter)
	newtondX, newtondIter := interpolationmethods.NewtonsDivided(x, y, xStar)
	fmt.Println("Newton's divided:", newtondX, "Iterations:", newtondIter)
	splineX, splineIter := interpolationmethods.Spline(x, y, xStar)
	fmt.Println("Spline:", splineX, "Iterations:", splineIter)
}
