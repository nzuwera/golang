package main

// import package
import (
	"fmt"
	"runtime"
	"reflect"
)

// Global variables
var currentOs = runtime.GOOS

// Main function
func main() {
	sayHello()
}

// User defined function
func sayHello() {
	// variable definition inside a function
	name := "Gilbert"
	fmt.Println("My name is ", name, " is Type of ", reflect.TypeOf(name))
	fmt.Println("Am using ", currentOs, " as OS, which is Type of ", reflect.TypeOf(currentOs))

	// Pointers = Memory address of a variable "name"
	namePointer := &name
	fmt.Println("This is the momery address of varialbe 'name'", namePointer, "and the value of 'name' is ", *namePointer)
}
