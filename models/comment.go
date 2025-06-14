package models

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primarykey;comment:'评论ID'" json:"id"`
	Content   string    `gorm:"type:text;not null;comment:'评论内容'" json:"content" binding:"required"`
	UserId    uint      `gorm:"not null;index;comment:'评论者ID'" json:"user_id"`
	PostId    uint      `gorm:"not null;index;comment:'评论所在文章ID'" json:"post_id"`
	CreatedAt time.Time `gorm:"comment:'创建时间'"`

	User User `gorm:"foreignkey:UserId;comment:'评论所属用户'" json:"user,omitempty"`
	Post Post `gorm:"foreignkey:PostId;comment:'评论所在文章'" json:"post,omitempty"`
}
