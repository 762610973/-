package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var x int32 = 100

func a() {
	atomic.AddInt32(&x, 1)
}
func main() {
	ticker := time.NewTicker(time.Second)
	/*count := 0
	for _ = range ticker.C {
		count++
		fmt.Println("ticker...")
		if count > 5 {
			ticker.Stop()
		}
	}*/
	chanInt := make(chan int)
	go func() {
		defer close(chanInt)
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Println("receive:", v)
		sum += v
		if sum >= 10 {
			break
		}
	}

}
