package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pleasedChord/gin-gorm-blog.git/models"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func NewCommentController(db *gorm.DB) *CommentController {
	return &CommentController{DB: db}
}

// 创建评论:已认证的用户可以对文章发表评论
func (ctrl *CommentController) CreateComment(c *gin.Context) {
	//校验入参
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	//校验用户
	var user models.User
	if err := ctrl.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录后评论"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//新增评论
	if err := ctrl.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论成功", "comment": comment})
}

// 读取评论:支持获取某篇文章的所有评论列表
func (ctrl *CommentController) GetComments(c *gin.Context) {
	var post models.Post

	postId := c.Param("post_id")
	if err := ctrl.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在"})
		return
	}

	var comments []models.Comment
	if err := ctrl.DB.Model(&models.Comment{}).
		Preload("User").Preload("Post").
		Where("post_id = ?", postId).
		Find(&comments).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "获取评论成功", "comments": comments})
}
