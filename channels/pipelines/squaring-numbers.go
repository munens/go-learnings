package main

import "fmt"

func main() {
	for v := range square(collect(2, 5, 6)) {
		fmt.Println(v) // 4, 25, 36
	}
}

func collect(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- (v * v)
		}
		close(out)
	}()

	return out
}
