package main

import (
	"fmt"
	"os"
)

func openCloseFile() {
	/*
		f, err := os.Open("onlyReadFile") //以只读方式打开文件
		defer f.Close()
		if err != nil {
			fmt.Println("err = ", err)
		} else {
			fmt.Println(f.Name())
		}*/
	//已读写方式打开文件，不存在则创建文件，权限是755
	f, err := os.OpenFile("test.md", os.O_RDWR|os.O_CREATE, 755)
	defer f.Close()
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(f.Name())
	}
}

//读操作
func readOps() {
	/*	f, err := os.Open("../fileNew.txt")
		for {
			if err != nil {
				fmt.Println("err = ", err)
			}
			buf := make([]byte, 3)
			n, err := f.Read(buf)
			if err == io.EOF {
				break
			}
			fmt.Println(n)
			fmt.Println(string(buf))
		}
		f.Close()*/
	/*	f, err := os.Open("../fileNew.txt")
		buf := make([]byte, 3)
		if err != nil {
			fmt.Println(err)
		}
		//从3开始读
		n, _ := f.ReadAt(buf, 3)
		fmt.Println(n)
		fmt.Println(string(buf))*/
	de, _ := os.ReadDir("F:\\StudyNotes\\Go\\标准库")
	for _, v := range de {
		fmt.Print(v.Name(), v.IsDir(), "\n")
	}

}

//定位
func seekFile() {
	f, err := os.Open("../fileNew.txt")
	if err != nil {
		fmt.Println(err)
	}
	/*b := make([]byte, 3)
	read, err := f.Read(b)
	fmt.Println(read, err)*/
	seek, err := f.Seek(-4, 2)
	fmt.Println(seek)
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 9)
	n, _ := f.Read(buf)
	fmt.Println("n = ", n)
	fmt.Println(string(buf))
	f.Close()
}

func main() {
	//openCloseFile()
	//readOps()
	seekFile()

}
