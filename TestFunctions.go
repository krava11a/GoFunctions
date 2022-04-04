package main

import (
	"fmt"
	"math"
)

func TestFunctions() {

	//Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	//Function closures
	fmt.Println("Function closures")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	//Exercise: Fibonacci closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	sum := -1
	f0 := 0
	f1 := 1
	return func() int {

		if sum == -1 {
			sum = 0

		} else {
			sum = f0 + f1
			f0 = f1
			f1 = sum
		}

		return sum
	}
}
