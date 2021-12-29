package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	words := strings.Split(str, " ")
	var resultWords []string
	for i := len(words) - 1; i >= 0; i-- {
		resultWords = append(resultWords, words[i])
	}
	var result string
	for _, word := range resultWords {
		result += word + " "
	}
	return result
}

func main() {
	str := "snow dog is god"
	str2 := reverse(str)
	fmt.Println(str2)
}
