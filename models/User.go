package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey;comment:'用户ID'" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique;not null;comment:'用户名'" json:"username" binding:"required"`
	Password  string    `gorm:"type:varchar(255);not null;comment:'用户密码'" json:"password"` //加密后的
	Email     string    `gorm:"type:varchar(255);unique;not null;comment:'用户邮箱'" json:"email"`
	CreatedAt time.Time `gorm:"comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'"`

	Posts    []Post    `gorm:"foreignkey:UserId;comment:'用户发布的文章'" json:"posts,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserId;comment:'用户发布的评论'" json:"Comments,omitempty"`
}

// db保存密码前，对密码做加密
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if tx.Statement.Changed("Password") {
		hashPassword, err := bcrypt.
			GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashPassword)
	}
	return nil
}

// 对比码是否匹配
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
