package main

import (
	"fmt"
	"reflect"
)

// related to 3 laws of reflection - https://blog.golang.org/laws-of-reflection

func main() {

	num := 9.2
	fmt.Printf("%T\n", num) // float64

	v := reflect.ValueOf(num)

	// Returns error - 'v' must be settable
	// v.SetFloat(9.5)

	fmt.Println("is v settable ? ", v.CanSet()) // is v settable ? false

	v2 := reflect.ValueOf(&num)
	fmt.Println("is v3 settable ? ", v2.CanSet()) // is v2 settable ? true
	v3 := v2.Elem()

	fmt.Println("is v3 settable ? ", v3.CanSet()) // is v3 settable ? true
	v3.SetFloat(9.5)

	fmt.Println("value of v3: ", v3.Interface()) // value of v3: 9.5
	fmt.Println("value of num: ", num) // value of num: 9.5

}
