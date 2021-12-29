package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	str := strings.ToLower(s)
	mapResult := make(map[rune]int)
	for _, symb := range str {
		mapResult[symb]++
	}
	if len(mapResult) != len(str) {
		return false
	}
	return true
}

func main() {
	str := "cd"
	fmt.Println("result is - ", check(str))
}
