package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm" // 注意导包
)

// 入门
// 注意gorm 的版本 当前v1.9 v2.x批量插入才友好
// 参考 https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func main() {
	// 连接数据库
	dsn := "root:000000@tcp(39.96.160.227:3306)/test?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{}) // 比如新增字段之后会自动新增到数据库
	// 创建数据行 创建实例
	// u2 := UserInfo{2, "王武1", 12, "男", "basketball"}
	// 更新全部对应字段
	//db.Create(&u2)
	// 更新部分对应字段
	// db.Select("name", "age", "gender").Create(&u2)
	// 上方操作取反 这些值不更新
	// db.Omit("Name", "Age", "id").Create(&u2)

	// 批量创建（）
	//users :=[]UserInfo{{Name: "z1" }, {Name: "z2" },{Name: "z2"}}
	//db.CreateInBatches(&users)
	//db.Model(&UserInfo{}).Create([]map[string]interface{}{
	//	{"Name": "jinzhu_1", "Age": 18},
	//	{"Name": "jinzhu_2", "Age": 20},
	//})
	//for _, user := range users {
	//	//fmt.Println("id is", user.ID)
	//}
	// 查询数据
	//var u UserInfo
	//db.First(&u) // 查询表中第一条数据
	//fmt.Println(u)
	//// 更新
	//db.Model(&u).Update("hobby", "双色球")
	//// 删除
	//db.Delete(&u)
}

// UserInfo 模型定义
type UserInfo struct {
	ID     uint
	Name   string
	Age    uint
	Gender string
	Hobby  string
}
