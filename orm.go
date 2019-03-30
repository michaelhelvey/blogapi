package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

// VerifyDBConnection prepares a connection to the database and panics
// if it can't connect.  Should be called at the very start of the API
func VerifyDBConnection() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "gustavus:^EH]4HhE[&C_HN=3@/blog?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
	return db
}

type postsQueryType struct {
	ID          int
	Title       string
	Body        string
	Author      int
	Created     time.Time
	UserID      int
	DisplayName string
	Avatar      string
	Email       string
	Username    string
}

/*
	GetPostsFromDB query needs to support optional
	1. searching by title
	2. filtering by tags, categories
	3. pagination
*/
func GetPostsFromDB(individual bool, postID string) ([]Post, error) {
	var posts []Post
	var err error
	var rows *sql.Rows
	if individual {
		rows, err = db.Query("select * from posts inner join users on users.id = posts.author where posts.id = ?", postID)
	} else {
		rows, err = db.Query("select * from posts inner join users on users.id = posts.author")
	}
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	var postsfromdb []postsQueryType
	for rows.Next() {
		var p postsQueryType
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.Author, &p.Created, &p.UserID, &p.DisplayName, &p.Avatar, &p.Email, &p.Username); err != nil {
			log.Fatalf("Could not scan post: %s", err.Error())
		}
		postsfromdb = append(postsfromdb, p)
	}

	// now we have a complete listing of the posts we need to processs, we can just
	// read those values into our front-facing values
	for _, p := range postsfromdb {
		a := User{p.UserID, p.DisplayName, p.Avatar, p.Email, p.Username}
		newP := Post{ID: p.ID, Title: p.Title, Body: p.Body, Author: a, Created: p.Created}
		posts = append(posts, newP)
	}
	return posts, err
}
