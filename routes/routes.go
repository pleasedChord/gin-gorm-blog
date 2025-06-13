package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pleasedChord/gin-gorm-blog.git/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	// 创建 Gin 引擎
	r := gin.Default()

	// 创建控制器实例
	userCtrl := controllers.NewUserController(db)

	public := r.Group("/api")
	{
		public.POST("/register", userCtrl.Resgister)
		public.POST("/login", userCtrl.Login)
	}

	return r
}
