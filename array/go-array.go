package main

import (
	"fmt"
)

func main() {

	//different ways for initializing array

	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	array3 := [3]int{1, 2}

	// if arrays have the same type
	// as example array1 , array2 and array3 have type [3]int - possible to compare arrays between each other

	if array1 == array2 {
		fmt.Println("array1 = ", array1)
		fmt.Println("array2 = ", array2)
		fmt.Println("array1 equals array2")
	}

	fmt.Println("---------------")

	if array2 != array3 {
		fmt.Println("array2 = ", array2)
		fmt.Println("array3 = ", array3)
		fmt.Println("array2 isn't equal array3")
	}

	fmt.Println("---------------")

	/*
		array4 := [4]int{1, 2}
		fmt.Println(array1 == array4) // will have problem during compilation int[3] != int[4]
	*/

}
