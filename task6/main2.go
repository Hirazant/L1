package main

import (
	"fmt"
	"time"
)

func someGoroutine2(chanel chan int) {
	for {
		fmt.Printf("do something with chanel\n")
	}
}

func main() {
	chanel := make(chan int)
	go someGoroutine2(chanel)

	time.Sleep(time.Second)
	close(chanel)
}
