package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main2() {

	// 给所有的表加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	// 连接数据库
	dsn := "root:000000@tcp(39.96.160.227:3306)/test?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) // 禁用复数 或者自己重写TableName方法 参考下方

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&Student{}) // 比如新增字段之后会自动新增到数据库

	// 创建table 方式2 这种有缺陷，不能创建第二次
	//db.Table("student_demo").CreateTable(&Student{})

	// 创建数据行
	//u1 := UserInfo{1, "张三", 12, "男", "football"}
	//db.Create(&u1)
}

// Student 注意mysql 5.6版本 index 707 字节限制
type Student struct {
	// gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64      // 修改成为 long
	DeletedAt *time.Time `sql:"index:dat"`

	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(10);unique_index:eml"` // 这个可以设置自己的名字
	Role         string  `gorm:"size:255"`                          // 设置字段大小为 255
	MemberNumber *string `gorm:"unique;not null;size:10"`           // 设置会员号 唯一键 不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                    // 设置num为自增类型
	Address      string  `gorm:"index:addr;size:30;default:'空地址'"`  // 为address 创建名为 addr 的索引
	IgnoreMe     int     `gorm:"-"`                                 // 忽略本字段
}
// 注意 sql.NullString
// 传值的时候 写 sql.NullString(String:"",Valid:true) 这个样代表我自己的空串有效果，不使用默认的


// 自己重命名
//func (Student)TableName() string {
//	return "student"
//}

// 以下是生成的mysql 的DDL语句
//CREATE TABLE `prefix_student` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,  // 自增主键
//`created_at` bigint(20) DEFAULT NULL, // long
//`updated_at` bigint(20) DEFAULT NULL,
//`deleted_at` datetime DEFAULT NULL,
//`name` varchar(255) DEFAULT NULL,
//`age` bigint(20) DEFAULT NULL,
//`birthday` datetime DEFAULT NULL,
//`email` varchar(10) DEFAULT NULL,
//`role` varchar(255) DEFAULT NULL,
//`member_number` varchar(10) NOT NULL, // 指定长度
//`num` int(11) DEFAULT NULL,
//`address` varchar(30) DEFAULT '空地址',
//PRIMARY KEY (`id`), // 指定主键
//UNIQUE KEY `member_number` (`member_number`),
//UNIQUE KEY `eml` (`email`), // 指定uk
//KEY `dat` (`deleted_at`),
//KEY `addr` (`address`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 // charset跟随数据库
