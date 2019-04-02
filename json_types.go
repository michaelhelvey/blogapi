package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// GormModel definition
type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// User is the type for users
type User struct {
	GormModel
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Username    string `json:"username"`
}

// Post is the type for blog posts in the API
type Post struct {
	GormModel
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	AuthorID   int        `json:"author_id"`
	Author     User       `gorm:"foreignkey:AuthorID" json:"author"`
	Tags       []Tag      `gorm:"many2many:posts_tags;" json:"tags"`
	Categories []Category `gorm:"many2many:posts_categories;" json:"categories"`
}

// Tag is the type for blog post tags
type Tag struct {
	GormModel
	Title string
}

// Category is the type for blog post categories
type Category struct {
	GormModel
	Title string
}

// InitOrm registers our models with beego's orm
func InitOrm() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	db, err = gorm.Open("mysql", "gustavus:^EH]4HhE[&C_HN=3@/blog?parseTime=true")
	if err != nil {
		return db, err
	}
	return db, err
}
