package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Удалить элемент по индексу i из a.

	// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(a[i:], a[i+1:])

	// 2. Усечь срез.
	a = a[:len(a)-1]

	fmt.Println(a) // [A B D E]
}
