package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// github.com/go-sql-driver/mysql  这个包 使用go get 的方式安装到了本地

// 数据库连接测试
func main() {
	// 主程序
	db, err := sql.Open("mysql", "root:teemo@tcp(127.0.0.1:3306)/beego_demo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("show tables")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		fmt.Println(rows.Columns())
	}

	fmt.Println("ok")
	// 避免编译出错的处理.
}
