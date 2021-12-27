package main

import "fmt"

type Human struct {
	pole1 int
	pole2 int
}

func (h *Human) method1() {
	fmt.Println("i'm Human!")
}

type Action struct {
	Human
}

func main() {
	var a Action
	a.method1()
}
