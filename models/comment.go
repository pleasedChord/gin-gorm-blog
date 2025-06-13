package models

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primarykey;comment:'评论ID'"`
	Content   string    `gorm:"type:text;not null;comment:'评论内容'"`
	UserId    uint      `gorm:"not null;index;comment:'评论者ID'"`
	PostId    uint      `gorm:"not null;index;comment:'评论所在文章ID'"`
	CreatedAt time.Time `gorm:"comment:'创建时间'"`

	User User `gorm:"foreignkey:UserId;comment:'评论所属用户'"`
	Post Post `gorm:"foreignkey:PostId;comment:'评论所在文章'"`
}
