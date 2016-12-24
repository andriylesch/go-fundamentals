package main

import "fmt"

func newIntOne() *int {
	return new(int)
}

func newIntTwo() *int {
	var dummy int
	return &dummy
}

func main() {

	// in current case variable1 has type *int which related to variable with type int
	variable1 := new(int)

	fmt.Println("")
	fmt.Print("variable1 address = ")
	fmt.Println(variable1)
	fmt.Println(" -------- ")

	fmt.Print("variable1 address = ")
	fmt.Println(*variable1)

	fmt.Println("-------- ")

	*variable1 = 10

	fmt.Print("variable1 changed data = ")
	fmt.Println(*variable1)

	// as exemple two functions newIntOne and newIntTwo has the same behavior

}
