package main

import (
	"time"
)

// User is the type for users
type User struct {
	ID          int
	DisplayName string
	Avatar      string
	Email       string
	Username    string
}

// Post is the type for blog posts in the API
type Post struct {
	ID         int
	Title      string
	Body       string
	Author     User
	Tags       []Tag
	Categories []Category
	Created    time.Time
}

// Tag is the type for blog post tags
type Tag struct {
	ID    int
	Title string
}

// Category is the type for blog post categories
type Category struct {
	ID    int
	Title string
}
