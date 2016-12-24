/*
	to check how it's working just add some parameters

	for example : go run main.go 1 2 3 4 5 "some text"

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	argsList := os.Args

	if len(argsList) == 0 {

	} else {
		for index, value := range argsList {
			fmt.Println("argsList[", index, "] = ", value)
		}
	}

}
