package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
)

var db *gorm.DB

// ServerErrorHandler Handles all server 500 errors
func ServerErrorHandler(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// NotFoundHandler returns not found responses
func NotFoundHandler(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusNotFound)
}

// PostsDetailHandler returns a JSON post by id
func PostsDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["post_id"])
	if err != nil {
		NotFoundHandler(w, "Invalid post id")
	}
	var post Post
	db.Preload("Author").First(&post, postID)

	jsonResponse, _ := json.Marshal(post)
	fmt.Fprintf(w, "%s", jsonResponse)
}

// PostsHandler returns an array of JSON posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	db.Preload("Author").Find(&posts)
	jsonResponse, _ := json.Marshal(posts)
	fmt.Fprintf(w, "%s", jsonResponse)
}

func main() {
	var err error
	db, err = InitOrm()
	if err != nil {
		panic("Cannot connect to the database")
	}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/posts/", PostsHandler)
	r.HandleFunc("/posts/{post_id}/", PostsDetailHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
