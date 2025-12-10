package main

import (
	"fmt"

	"github.com/fernoe1/CM/assignmentone/iterativemethods"
)

func main() {
	expr := "exp(x) - x**2"
	dExpr := "exp(x) - 2*x"
	gExpr := "-sqrt(exp(x))"

	x0 := float64(-2)
	x1 := float64(1)
	x2 := -1.0

	tol := 1e-6
	maxN := 1000

	rootBis, _, iterBis := iterativemethods.Bisection(expr, x0, x1, tol, maxN)
	fmt.Printf("Bisection: %d iterations. Root: %f\n", iterBis, rootBis)

	rootFix, _, iterFix := iterativemethods.FixedPoint(gExpr, x0, tol)
	fmt.Printf("FixedPoint: %d iterations. Root: %f\n", iterFix, rootFix)

	rootNewton, _, iterNewton := iterativemethods.NewtonRaphson(expr, dExpr, x0, tol)
	fmt.Printf("Newton: %d iterations. Root: %f\n", iterNewton, rootNewton)

	rootSec, _, iterSec := iterativemethods.Secant(expr, x0, x1, tol, maxN)
	fmt.Printf("Secant: %d iterations. Root: %f\n", iterSec, rootSec)

	rootFalse, _, iterFalse := iterativemethods.FalsePosition(expr, x0, x1, tol, maxN)
	fmt.Printf("False: %d iterations. Root: %f\n", iterFalse, rootFalse)

	rootMuller, _, iterMuller := iterativemethods.Muller(expr, x0, x1, x2, tol, maxN)
	fmt.Printf("Muller %d iterations. Root: %f\n", iterMuller, rootMuller)

	//fmt.Println(iterativemethods.Bisection("x**3 - x - 2", 1, 2, 1e-3, 13))
	//
	//fmt.Println(iterativemethods.FixedPoint("cos(x)", 0.5, 1e-6))
	//
	//fmt.Println(iterativemethods.NewtonRaphson("x**3 - 2*x - 5", "3*x**2 - 2", 2, 1e-7))
	//
	//fmt.Println(iterativemethods.Secant("x**3 - 2*x - 5", 2, 3, 1e-6, 1000))
	//
	//fmt.Println(iterativemethods.FalsePosition("2 * exp(x) * sin(x) - 3", 0, 1, 1e-6, 1000))
	//
	//fmt.Println(iterativemethods.Muller("cos(x) - x * exp(x)", -1, 0, 1, 1e-6, 10000))
}
