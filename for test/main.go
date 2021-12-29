package main

import (
	"fmt"
	"unsafe"
)

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]

	fmt.Println("v = ", v)
	fmt.Println("justString = ", justString)
	fmt.Println("len(v) = ", len(v))
	fmt.Println("len(justString) = ", len(justString))
	fmt.Println("cap(v) = ", cap([]byte(v)))
	fmt.Println("cap(justString) = ", cap([]byte(justString)))
	fmt.Println("sizeof(v) = ", unsafe.Sizeof(v))
	fmt.Println("sizeof(justString) = ", unsafe.Sizeof(justString))
}

func main() {
	someFunc()
}
