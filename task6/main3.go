package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func waitGoroutine() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			doSomethingGoroutine(id)
			wg.Done()
		}(i)
	}
	//Останавливаем горутину и ждем выполнения предыдущих
	wg.Wait()
	fmt.Printf("All Done!")
}

func doSomethingGoroutine(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("GoRoutine %d, do something\n", id)
	}
}

func main() {
	waitGoroutine()
}
