package main

import "fmt"

func main() {
	//Assigning function to the variable.
	var (
		area = func(l int, b int) int {
			return l * b
		}
	)
	fmt.Println(area(20, 30))

	//Passing arguments to anonymous functions.
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	//Closures Functions
	//Anonymous function accessing variable
	//on each iteration of loop inside function body.
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}

	//Higher Order Functions
	//Passing Functions as Arguments to other Functions
	partial := partialSum(3)
	fmt.Println(partial(7))

	//Returning Functions from other Functions
	fmt.Println(squareSum(5)(6)(7))

	//User Defined Function Types
	fmt.Println(squareSumUserDefine(5)(6)(7))
}

func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

type First func(int) int
type Second func(int) First

func squareSumUserDefine(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}
