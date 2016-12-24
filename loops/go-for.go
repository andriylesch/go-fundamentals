package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("\n standard loop")
	for index := 0; index < len(array); index++ {
		fmt.Print(" ", array[index])
	}

	fmt.Println("\n foreach loop (range)")
	// for index, item := range collection
	for _, item := range array {
		fmt.Print(" ", item)
	}

	fmt.Println("\n infinity loop ")

	index := 0
	for {
		if index >= len(array) {
			break
		}

		fmt.Print(" ", array[index])
		index++
	}

	fmt.Println("\n while loop ")
	i := 0
	for i < 10 {

		fmt.Print(" ", i)
		i++
	}

}
