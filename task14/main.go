package main

import "fmt"

func typ(data interface{}) {
	fmt.Printf("%v is %T\n", data, data)
}

func main() {
	typ("str")
	typ(33)
	typ(33.33)
	typ(make(chan int))
}
