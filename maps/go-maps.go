package main

import (
	"fmt"
)


/*

    Maps in Go language - sometimes called hash-table or dictionary in other languages

    dict := make(map[key-type]value-type)


*/

func main() {


    // dictionary with data
    dict := map[string]int {
        "key1" : 11,
        "key2" : 22,
        "key3" : 33 
    } 

    fmt.Println("dict = ", dict)

    // dictionary equivalent to to previous one
    dict1 := make(map[string]int)

    //adding elements
    dict1["key1"] = 11
    dict1["key2"] = 22
    dict1["key3"] = 33

    fmt.Println("dict1 ", dict1)

    //change element
    dict1["key2"] = 55
    fmt.Println("dict1 change element", dict1)


    //remove element
    delete(dict1,"key3")
    fmt.Println("dict1 remove element", dict1)

}