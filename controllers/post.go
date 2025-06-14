package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pleasedChord/gin-gorm-blog.git/models"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

// 创建文章
func (ctrl *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章创建成功", "postId": post.ID})
}

// 根据id读取文章
func (ctrl *PostController) GetPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := ctrl.DB.Preload("User").Preload("Comments.User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// 读取文章列表
func (ctrl *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := ctrl.DB.Model(&models.Post{}).Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// 更新文章
func (ctrl *PostController) UpdatePost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	if err := ctrl.DB.Find(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}

	//校验当前用户是否与文章用户一致
	userId, exits := c.Get("user_id")
	if !exits {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少参数"})
		return
	}

	if userId != post.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有本人可以修改此文章"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "修改成功", "post": post})
}

// 删除章
func (ctrl *PostController) DeletePost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")
	if err := ctrl.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少参数"})
		return
	}
	if userId != post.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限"})
		return
	}

	if err := ctrl.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功删除"})
}
