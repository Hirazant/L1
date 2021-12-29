package main

import (
	"fmt"
	"strings"
)

func check2(s string) bool {
	str := strings.ToLower(s)
	runestr := []rune(str)
	mapResult := make(map[rune]int)
	for _, symb := range runestr {
		mapResult[symb]++
	}
	if len(mapResult) != len(runestr) {
		return false
	}
	return true
}

func main() {
	str := "ау"
	fmt.Println("result is - ", check2(str))
}
