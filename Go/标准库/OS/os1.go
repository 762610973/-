package main

//os包提供了平台无关性的一些功能

import (
	"fmt"
	"os"
)

func createFile() {
	//重新执行会覆盖
	file, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println("err :", err)
	} else {
		fmt.Println("文件名为：", file.Name())
	}
}
func makeDir() {
	err := os.Mkdir("dir1.txt", os.ModePerm)
	if err != nil {
		fmt.Println("err = ", err)
	}
	err1 := os.MkdirAll("a/b/c", os.ModePerm)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
	}
}
func remove() {
	//删除不存在的文件会报错
	err := os.Remove("file1.txt")
	if err != nil {
		fmt.Println("err = ", err)
	}
	err = os.RemoveAll("a")
	if err != nil {
		fmt.Println("err = ", err)
	}
}
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("当前工作目录是:", dir)
	}
	os.Chdir("e:/")
	dir, _ = os.Getwd()
	fmt.Println("dir = ", dir)
	s := os.TempDir()
	fmt.Println("当前临时目录是:", s)
}

func reName() {
	err := os.Rename("file1.txt", "fileNew.txt")
	if err != nil {
		fmt.Println("err = ", err)
	}
}

func readFile() {
	file, err := os.ReadFile("F:\\StudyNotes\\Go\\面试.md")
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println(string(file))
	}
}
func writeFile() {
	os.WriteFile("fileNew.txt", []byte("hello\nhello"), os.ModePerm)

}
func main() {
	a := []int{1}
	a = a[1:]
	fmt.Println(a)
	//createFile()
	//makeDir()
	//remove()
	//getWd()
	//reName()
	//readFile()
	//writeFile()

}
