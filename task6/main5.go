package main

import (
	"fmt"
	"time"
)

func someGoroutine3(flag *bool) {
	for {
		if *flag {
			fmt.Println("it's all")
			return
		}
	}
}

func main() {
	flag := false
	go someGoroutine3(&flag)

	time.Sleep(time.Second)
	flag = true
	time.Sleep(time.Second)
}
