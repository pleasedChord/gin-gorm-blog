package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pleasedChord/gin-gorm-blog.git/models"
	"github.com/pleasedChord/gin-gorm-blog.git/util"
	"gorm.io/gorm"
)

// 受限接口鉴权
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHead := c.GetHeader("Authorization")
		if authHead == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求非法"})
			c.Abort()
		}

		autherHeadSplit := strings.Split(authHead, " ")
		if len(autherHeadSplit) != 2 && autherHeadSplit[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求非法"})
			c.Abort()
		}

		tokenString := autherHeadSplit[1]
		claims, err := util.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token"})
			c.Abort()
		}

		var user models.User
		if db.Model(&models.User{}).Where("user_id = ?", claims.UserId).First(&user).Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			c.Abort()
		}

		user.Password = ""
		c.Set("user_id", user.ID)

		c.Next()
	}
}
