package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		file, _ := os.ReadFile("./test.html")
		//writer.Write([]byte("测试ing"))
		writer.Write(file)
		writer.Write(file)
	})
	http.ListenAndServe(":8080", nil)
	/*res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Header)
	fmt.Println(res.Status)*/

}
