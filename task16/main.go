package main

import (
	"fmt"
	"time"
)

type incr struct {
	i int
}

func (i *incr) increment() {
	i.i = i.i + 1
}

func work(inc *incr) {
	inc.increment()
}

func main() {
	inc := incr{0}
	for i := 0; i < 5; i++ {
		go work(&inc)
	}
	time.Sleep(time.Second)

	fmt.Println(inc)
}
