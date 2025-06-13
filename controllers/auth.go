package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pleasedChord/gin-gorm-blog.git/models"
	"github.com/pleasedChord/gin-gorm-blog.git/util"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// 注入
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// 用户注册
func (ctrl *UserController) Resgister(c *gin.Context) {
	//绑定请求数据
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//校验入参
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名、密码和邮箱都不能为空"})
		return
	}

	//检查数据库是否已经有用户名和邮箱
	var count int64
	result := ctrl.DB.Model(&models.User{}).
		Where("username = ? or email = ?", user.Username, user.Email).
		Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或邮箱已存在"})
		return
	}

	if err := ctrl.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户创建失败" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "用户创建成功",
		"user": gin.H{
			"userId":   user.ID,
			"userName": user.Username,
			"email":    user.Email,
		},
	})
}

// 用户登录
func (ctrl *UserController) Login(c *gin.Context) {
	//获取入参
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//校验入参
	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "必须输入用户名和密码"})
		return
	}

	var user models.User
	if err := ctrl.DB.Model(&models.User{}).
		Where("username = ?", request.Username).
		First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "账号不存在"})
		return
	}

	//验证密码
	if !user.ComparePassword(request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "账号密码不正确"})
	}

	//使用JWT生成token
	token, err := util.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "生成token失败"})
		return
	}

	//返回
	c.JSON(http.StatusCreated, gin.H{
		"message": "登录成功",
		"token":   token,
		"user": gin.H{
			"userId":   user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
