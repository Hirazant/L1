package main

import "fmt"

type kilovolts int

func toKilovolts(data int) kilovolts {
	return kilovolts(data * 1000)
}

func something(v kilovolts) {
	fmt.Printf("value - %d, type - %T", v, v)
}

func main() {
	volts := 25
	something(toKilovolts(volts))
}
