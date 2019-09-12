package main

// import package
import (
	"fmt"
	"math"
)

// Main function
func main() {

	/* local variable definition */
	var a int = 100
	var b int = 200
	var ret int

	/* calling a function to get max value */
	ret = max(a, b)
	fmt.Println("Max value is : ", ret)

	/* calling a swap function */
	c, d := swap("Mahesh", "Kumar")
	fmt.Println(c, d)

	/*
		03 - Declare a function as variable
	*/

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	/*
		Function Methods
	*/
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())

}

/*
	01 - Calling a Function:
	function returning the max between two numbers
*/
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
	02 - Returning multiple values from Function:
	function which swap two given string
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
	Function Methods
*/
/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
