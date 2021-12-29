package main

import "fmt"

func main() {
	first, second := 22, 11

	first, second = second, first
	fmt.Printf("first - %d, second - %d\n", first, second)

	first = first + second
	second = first - second
	first = first - second
	fmt.Printf("first - %d, second - %d", first, second)
}
