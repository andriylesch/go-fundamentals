package main

import "fmt"

/*

   Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
   Slices don't have limitation of items as array, during initialization of slice length = 0, capacity = 0.
   It's possible to say that slice is light version of array.

   good link - https://blog.golang.org/go-slices-usage-and-internals


*/

func main() {

	fmt.Println("---- Slices ------")

	// init slice
	slice := make([]string, 3) // in C# it will be like List<string>

	fmt.Println("slice empty:", slice)

	// different ways to add element in slice
	slice[0] = "text1"

	slice = append(slice, "text2")
	slice = append(slice, "text3")
	slice = append(slice, "text4")

	fmt.Println("slice added elements:", slice)

	// change data by index
	slice[0] = "some text"

	fmt.Println("slice changed item:", slice)

	// remove element from slice
	slice[0] = slice[len(slice)-1]
	slice = slice[:len(slice)-1]

	fmt.Println("slice removed item:", slice)

	// copy slice
	slicecopy := make([]string, len(slice))
	copy(slicecopy, slice)

	fmt.Println("slicecopy:", slicecopy)

	// get length and capacity of slice

	fmt.Println("slice length :", len(slice))
	fmt.Println("slice capacity :", cap(slice))

}
