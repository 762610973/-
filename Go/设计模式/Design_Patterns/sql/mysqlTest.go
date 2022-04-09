package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mytest2"
	db, err := sql.Open("mysql", dsn) //open并没有真正的发起连接，只是检查格式
	//defer db.Close()//?写在这里是不行的
	if err != nil {
		panic(err)
	}
	//做完错误检查之后，确保db不为nil
	defer db.Close() //释放资源
	err = db.Ping()  //发起连接
	if err != nil {
		fmt.Println("连接失败,err:", err.Error())
		return
	}

	fmt.Println("success")
}
