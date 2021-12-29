package main

import "fmt"

func toX2(secondChan chan<- int, firstChan <-chan int) {
	for val := range firstChan {
		secondChan <- val * 2
	}
}

func main() {
	mas := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	firstChan := make(chan int)
	secondChan := make(chan int, len(mas))
	defer close(secondChan)

	go toX2(firstChan, secondChan)

	for _, val := range mas {
		firstChan <- val
	}

	close(firstChan)

	for _, _ = range mas {
		fmt.Println(<-secondChan)
	}

}
