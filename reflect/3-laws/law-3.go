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

	fmt.Println("is v settable ?: ", v.CanSet())

	v2 := reflect.ValueOf(*num)
	v3 := v2.Elem()

	fmt.Println("is v3 settable ?: ", v.CanSet())
	v3.SetFloat(9.5)

	fmt.Println("value of v3: ", v3.Interface())
	fmt.Println("value of num: ", num)

}
