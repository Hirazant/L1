package main

import "fmt"

func kilo(data int) int {
	return data * 1000
}

func main() {
	volts := 25
	fmt.Println("volts is - ", volts, ", but kilovolts is - ", kilo(volts))
}
