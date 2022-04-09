package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var wg sync.WaitGroup
var i = 0

func add() {
	lock.Lock()
	defer wg.Done()
	i++
	fmt.Println(i)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	defer wg.Done()
	i--
	fmt.Println(i)
	lock.Unlock()
}

func main() {
	//runtime.GOMAXPROCS(4)
	for i := 0; i < 25; i++ {
		wg.Add(1)
		add()
		wg.Add(1)
		sub()
	}
	wg.Wait()
	fmt.Println("end i", i)
	fmt.Println("end")
}
