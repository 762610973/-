package main

import (
	"fmt"
	"time"
)

/*
var c = make(chan int)

func main() {
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}()
	//	for v := range c {
	//	fmt.Println(v)
	//}

	//for {
	//	v, ok := <-c
	//	if ok {
	//		fmt.Println(v)
	//	} else {
	//		break
	//	}
	//}

	//	r := <-c
	//	fmt.Println(r)
	//	r = <-c
	//	fmt.Println(r)
	//	r = <-c
	//	fmt.Println(r)
	//通道没有关闭，再次读就会死锁
	//必须关闭才行，关闭之后再读就是0
}*/
func main() { /*
		timer1 := time.NewTimer(time.Second * 2)
		fmt.Println(time.Now())
		t1 := <-timer1.C
		fmt.Println(t1)*/
	/*fmt.Println(time.Now())
	<-time.After(time.Second * 3)
	fmt.Println(time.Now())*/
	/*timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("func...")
	}()

	s := timer3.Stop()
	if s {
		fmt.Println("stop...")
	}
	time.Sleep(time.Second * 2)
	fmt.Println("end")*/
	fmt.Println("before")
	timer4 := time.NewTimer(time.Second * 5)
	timer4.Reset(time.Second * 2)
	//_ = <-timer4.C
	<-timer4.C
	fmt.Println("after")
}
