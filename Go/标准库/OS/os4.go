package main

import (
	"fmt"
	"os"
)

func main() {
	/*	fmt.Println(os.Getegid())
		fmt.Println(os.Getppid())
		p2, _ := os.FindProcess(-1)
		fmt.Println(p2)*/
	/*s := os.Environ()
	fmt.Println(s)*/
	s2 := os.Getenv("GOPATH")
	fmt.Println(s2)
	s, b := os.LookupEnv("GOPATH")
	if b {
		fmt.Println(s)
	}
	//替换
	os.Setenv("env1", "env_")
	fmt.Println(os.Getenv("env1"))
	// os.Clearenv()
}
