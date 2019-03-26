package main

import (
	"math/rand"
	"time"
)

// TestPostBinary is a struct containing a Post struct instance and
// the byte string which that post serializes into, ultimately representing
// the source and hoped-for value from JSON serializing the Post.
type TestPostBinary struct {
	in   Post
	want []byte
}

// MockUser creates a User instance with example data
func MockUser() User {
	userID := rand.Int()
	userResult := User{userID, "Michael Helvey", "https://example.com/jpg", "test@whatever.com", "helveticus"}
	return userResult
}

// MockPost creates a Post instance with example data
func MockPost() Post {
	postID := rand.Int()
	createdTime := time.Now()

	user := MockUser()
	post := Post{postID, "Test Title", "Lorum Ipsum etc etc", user, createdTime}
	return post
}
