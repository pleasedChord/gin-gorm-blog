package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pleasedChord/gin-gorm-blog.git/controllers"
	"github.com/pleasedChord/gin-gorm-blog.git/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	// 创建 Gin 引擎
	r := gin.Default()

	// 创建控制器实例
	userCtrl := controllers.NewUserController(db)
	postCtrl := controllers.NewPostController(db)
	commentCtrl := controllers.NewCommentController(db)

	public := r.Group("/api")
	{
		//用户注册和登录
		public.POST("/register", userCtrl.Register)
		public.POST("/login", userCtrl.Login)

		//文章查看
		public.GET("/post/getPosts", postCtrl.GetPosts)
		public.GET("/post/post/:id", postCtrl.GetPost)

		//评论查看
		public.GET("/comment/getComments/:post_id", commentCtrl.GetComments)
	}

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware(db))
	{
		auth.GET("/user/me", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(http.StatusOK, gin.H{"user": user})
		})

		//文章管理
		auth.POST("/post/create", postCtrl.CreatePost)
		auth.PUT("/post/updatePost/:id", postCtrl.UpdatePost)
		auth.DELETE("/post/deletePost/:id", postCtrl.DeletePost)

		//评论管理
		auth.GET("/comment/create", commentCtrl.CreateComment)
	}

	return r
}
