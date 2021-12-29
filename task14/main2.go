package main

import "fmt"

func identifier(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("it's int!")
	case string:
		fmt.Println("it's string!")
	case float64:
		fmt.Println("it's float64!")
	case rune:
		fmt.Println("it's rune!")
	case bool:
		fmt.Println("it's bool!")
	case chan int:
		fmt.Println("it's chan int!")
	}
}

func main() {
	var data []interface{}
	data = append(data, "strrr")
	data = append(data, 233)
	data = append(data, 'd')
	data = append(data, true)
	data = append(data, 34.4)
	data = append(data, make(chan int))

	for _, val := range data {
		identifier(val)
	}

}
