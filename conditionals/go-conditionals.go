package main

import (
	"fmt"
	//"os"
	"time"
)

type CustomType struct {
	Name string
}

func main() {

	// if condition
	firstVal, secondVal := 11, 33

	//base if
	if firstVal < secondVal {
		fmt.Println("firstVal < secondVal")
	}

	//embedded statement
	if fVal, sVal := 15, 30; fVal > sVal {
		fmt.Println("fVal > sVal")
	} else if fVal < sVal {
		fmt.Println("fVal < sVal")
	} else {
		fmt.Println("other condition")
	}

	// embedded statement
	fmt.Println("-------------------------")
	switch i := 2; i {
	case 1, 2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default value")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("it's working day")
	default:
		fmt.Println("it's something else")
	}

	//switch on nothing

	name := "Andriy"

	switch {
	case name == "Andriy", name == "Vitaliy":
		fmt.Println("it's Andriy or Vitaliy")
	case name == "Maxim":
		fmt.Println("it's Maxim")
	default:
		fmt.Println("I don't know")
	}

	SwitchType("string type")
	SwitchType(123)
	SwitchType(CustomType{Name: "CustomType"})

}

func SwitchType(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("it's int")
	case string:
		fmt.Println("it's string")
	case CustomType:
		fmt.Println("it's CustomType")
	default:
		fmt.Println("unknown type")
	}
}
