package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{9, 16}

	s1 := square(collect(numbers))
	s2 := square(collect(numbers))

	for v := range merge(s1, s2) {
		fmt.Println("final:", v)
	}

}

func collect(in []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range in {
			fmt.Println("to add", v)
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
			val := v
			fmt.Println("to square: ", val)
			out <- val * val
		}

		close(out)
	}()

	return out
}

func merge(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(ins))
	// fan out:
	for _, v := range ins {
		go func() {
			// fan in:
			val := <-v
			fmt.Println("to merge:", val)
			out <- val
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
