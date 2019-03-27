package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

// VerifyDBConnection prepares a connection to the database and panics
// if it can't connect.  Should be called at the very start of the API
func VerifyDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "gustavus:^EH]4HhE[&C_HN=3@localhost/blog")
	if err != nil {
		panic("Could not open connection to sql database")
	}
	if err := db.Ping(); err != nil {
		panic("Could not open connection to sql database")
	}
	return db
}

/*
	Posts query needs to support optional
	1. searching by title
	2. filtering by tags, categories
	3. pagination
*/
func GetPostsFromDB() ([]Post, error) {
	var posts []Post
	var err error
	rows, err := db.Query("select * from posts")
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Body, &p.Created); err != nil {
			log.Fatalf("Could not scan post with id %d", p.ID)
		}
		posts = append(posts, p)
	}
	return posts, err
}
