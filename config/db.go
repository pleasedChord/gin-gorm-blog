package config

import (
	"fmt"
	"log"

	"github.com/pleasedChord/gin-gorm-blog.git/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("连接db失败,err=", err)
		return nil
	}
	fmt.Println("连接db成功")

	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalln("数据库迁移失败,err=", err)
		return nil
	}
	fmt.Println("数据库迁移成功")
	return db
}
