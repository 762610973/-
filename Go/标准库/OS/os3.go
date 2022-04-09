package main

import "os"

func write() {
	name := "test.md"
	f, _ := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0777)
	defer f.Close()
	f.Write([]byte("# hello world\n"))
}

func writeString() {
	f, _ := os.OpenFile("../fileNew.txt", os.O_RDWR|os.O_APPEND, 777)
	defer f.Close()
	f.WriteString("hello golang")

}
func writeAt() {
	//在某一个索引位置开始写
	//会覆盖后面的内容
	f, _ := os.OpenFile("../fileNew.txt", os.O_RDWR, 0777)
	f.WriteAt(
		[]byte("1234"),
		6,
	)
	f.Close()
}
func main() {
	//write()
	//writeString()
	writeAt()
}
