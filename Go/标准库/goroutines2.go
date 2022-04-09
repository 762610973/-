package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string) {
	fmt.Println("step1", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("step2", url)
	defer response.Body.Close()
	fmt.Println("step3", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("step4:", len(body))
}
func f0() (i int) { //6
	i = 5
	defer func() {
		i++
	}()
	return
}
func f() int { //5
	i := 5
	defer func() {
		i++
	}()
	return i
}
func f1() (result int) { //5
	t := 5
	defer func() {
		t += 5
	}()
	return t
}
func f2() (r int) { //5
	t := 5
	defer func() {
		t += 5
	}()
	return t
}
func f3() (r int) { //1
	defer func(r int) {
		r += 5
	}(r)
	return 1
}
func f4() {
	defer func() {

		defer func() {

			fmt.Println("1")
		}()
		fmt.Println("外")
	}()
}
func main() {
	/*	go responseSize("https://kamier.top")
		go responseSize("https://baidu.com")
		go responseSize("https://qq.com")
		time.Sleep(10 * time.Second)*/
	fmt.Println(f0())
	fmt.Println(f())
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	f4()
}
