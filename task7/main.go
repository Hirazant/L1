package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type chanMap struct {
	key   int
	value int
}

func toData(ch chan chanMap, data map[int]int, sn *sync.Mutex) {
	c := <-ch
	sn.Lock()
	data[c.key] = c.value
	sn.Unlock()
}

func toChanMap(ch chan chanMap) {
	key := rand.Intn(100)
	value := rand.Intn(100)
	ch <- chanMap{key, value}
}

func main() {
	ch := make(chan chanMap)
	data := make(map[int]int)
	sn := sync.Mutex{}
	defer close(ch)
	for i := 0; i < 10; i++ {
		go toChanMap(ch)
		go toData(ch, data, &sn)
	}

	time.Sleep(time.Second)
	fmt.Println(data)
}
