package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/pleasedChord/gin-gorm-blog.git/config"
	"github.com/pleasedChord/gin-gorm-blog.git/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()

	// 设置路由
	r := routes.SetupRoutes(db)

	port := config.GetEnv("PORT", "8080")
	log.Printf("服务器启动在 http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
