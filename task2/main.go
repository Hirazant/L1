package main

import (
	"fmt"
	"sync"
)

func square(c int) {
	fmt.Println(c * c)
}

func main() {
	a := [4]int{3, 4, 6, 7}
	wg := sync.WaitGroup{}
	for _, v := range a {
		wg.Add(1)
		go func() {
			square(v)
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println("it's all")

}
