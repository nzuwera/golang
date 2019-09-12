package main

// import package
import "fmt"

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
