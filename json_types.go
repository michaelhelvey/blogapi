package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// PostsResponse represents the JSON serialized response for posts queries
type PostsResponse struct {
	data []Post
}

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

/* REMOVE */

// ToJSON creates a JSON byte string from a Post struct instance
func (p Post) ToJSON() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(fmt.Sprintf("Post %x could not be serialized", p.ID))
	}
	return b
}

// ToJSON creates a JSON byte string from a Post struct instance
func (p PostsResponse) ToJSON() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(fmt.Sprintf("Posts %x could not be serialized"))
	}
	return b
}

// PostFromJSON takes in a byte string and creates a Post instance
func PostFromJSON(b []byte) (Post, error) {
	var p Post
	err := json.Unmarshal(b, &p)
	if err != nil {
		// we don't want to panic, we just want to return the error so
		// that a HTTP handler can handle it or something
		return Post{}, err
	}
	return p, nil
}
