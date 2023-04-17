package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

//	func main() {
//		db, err := gorm.Open(mysql.New(mysql.Config{
//			DSN:                       "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
//			SkipInitializeWithVersion: true,                                                                            // 根据当前 MySQL 版本自动配置
//		}), &gorm.Config{})
//		if err != nil {
//			panic("connection panic")
//		}
//		user := User{Name: "fzc", Age: 18, Birthday: time.Now()}
//
//		result := db.Create(&user) // 通过数据的指针来创建
//
//		fmt.Printf("%v %v %v", result.Error, user.ID, result.RowsAffected)
func main() {
	for i := 0; i < 10; i++ {
		id := uuid.New()
		fmt.Printf("%s %s\n", id, id.Version().String())
	}

	for i := 0; i < 10; i++ {
		id2, err := uuid.NewRandom()
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%s %s\n", id2, id2.Version().String())
	}

}
