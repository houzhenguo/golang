package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"  // 注意别缺了这些引用
	"github.com/jmoiron/sqlx"
)

func main() {
	initDB()
	fmt.Println("init db done", db)
	queryRowDemo()
}
var db *sqlx.DB

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

func initDB() (err error) {
	dsn := "root:000000@tcp(39.96.160.227:3306)/test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
type user struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}