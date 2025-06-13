package models

import "time"

type Post struct {
	ID        uint      `gorm:"primarykey;comment:'文章ID'"`
	Title     string    `gorm:"type:varchar(255);not null;comment:'文章标题'"`
	Content   string    `gorm:"type:text;not null;comment:'文章内容'"`
	UserId    uint      `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'"`

	User     User      `gorm:"foreignkey:UserId;comment:'文章作者'"`
	Comments []Comment `gorm:"foreignkey:PostId;comment:'文章的评论'"`
}
