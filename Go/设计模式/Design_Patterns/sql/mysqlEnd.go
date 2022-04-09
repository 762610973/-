package main

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

func initMysql() (err error) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/mytest2"
	//初始化全局的db对象而不是声明一个db变量
	db, err = sql.Open("mysql", dsn) //open并没有真正的发起连接，只是检查格式

	if err != nil {
		panic(err)
	}
	//做完错误检查之后，确保db不为nil
	defer db.Close() //释放资源
	err = db.Ping()  //发起连接
	if err != nil {
		fmt.Println("连接失败,err:", err.Error())
		return err
	}
	db.SetConnMaxLifetime(time.Second * 1000) //存活最长时间,根据业务来做
	db.SetMaxOpenConns(200)                   //最大连接数
	db.SetMaxIdleConns(10)                    //最大空闲连接数
	return

}
func main() {
	if err := initMysql(); err != nil {
		fmt.Println("connect to db failed,err:", err)
	}
	defer db.Close()
	fmt.Println("success connect to db")
}
