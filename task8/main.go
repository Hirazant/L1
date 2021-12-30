package main

import "fmt"

func main() {
	var number int64 = 100
	var numberOfBite int64 = 5
	switchToOne := false
	switch switchToOne {
	case true:
		number |= 1 << numberOfBite
	default:
		number &^= 1 << numberOfBite
	}

	fmt.Println("result is - ", number)
}
