package main

import (
	"fmt"
	"time"
)

func someGoroutine(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Printf("not work\n")
			return
		default:
			fmt.Printf("work\n")
		}
	}
}

func main() {
	quit := make(chan bool)
	defer close(quit)
	go someGoroutine(quit)

	time.Sleep(time.Second)
	quit <- true
}
