package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seconds = 10

func readFromChanel(chanel chan int) {

	for data := range chanel {
		fmt.Printf("now data is %d\n", data)
	}
}

func writeToChanel(chanel chan int) {
	for {
		data := rand.Intn(100)
		chanel <- data
	}
}

func main() {
	chanel := make(chan int)
	defer close(chanel)
	go writeToChanel(chanel)
	go readFromChanel(chanel)

	time.Sleep(time.Duration(seconds) * time.Second)
}