package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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
	postID := vars["post_id"]
	posts, err := GetPostsFromDB(true, postID)
	if err != nil {
		ServerErrorHandler(w, err)
		return
	}
	if len(posts) == 0 {
		NotFoundHandler(w, "Post Not Found")
		return
	}
	postsResponseJSON, err := json.Marshal(posts[0])
	fmt.Fprintf(w, "%s", postsResponseJSON)
}

// PostsHandler returns an array of JSON posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPostsFromDB(false, "_")
	if err != nil {
		ServerErrorHandler(w, err)
		return
	}
	postsResponseJSON, err := json.Marshal(posts)
	if err != nil {
		ServerErrorHandler(w, err)
		return
	}
	fmt.Fprintf(w, "%s", postsResponseJSON)
}

func main() {
	// verify that we can connect to the db
	VerifyDBConnection()
	// if we can, great, create our API
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/posts/", PostsHandler)
	r.HandleFunc("/posts/{post_id}/", PostsDetailHandler)
	// TODO:
	// search by title
	// categories
	// tags
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
