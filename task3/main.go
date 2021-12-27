package main

import (
	"fmt"
	"sync"
)

func square(k int) {
	sum += k * k
}

var sum int

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, val := range a {
		wg.Add(1)

		go func(val int) {
			square(val)
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println("sum is - ", sum)
}
