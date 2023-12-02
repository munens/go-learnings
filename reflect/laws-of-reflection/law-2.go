package main

import (
	"os"
	"io"
	"fmt"
	"reflect"
)

// related to 3 laws of reflection - https://blog.golang.org/laws-of-reflection

func main() {

	var r io.Reader
	tty, _ := os.OpenFile("../channels/pipelines/fan-out-fan-in.go", os.O_RDWR, 0)

	r = tty
	fmt.Printf("%T\n", r) // *os.File
	fmt.Println("r: ", r) // &{0xc0000a2120}

	v := reflect.ValueOf(r)
	fmt.Printf("%T\n", v) // reflect.Value
	fmt.Println("v: ", v) // &{0xc0000a2120}

	w := v.Interface().(io.Writer)
	fmt.Printf("%T\n", w) // *os.File
	fmt.Println("w: ", w) // &{0xc0000a2120}

}
