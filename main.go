package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// PostsDetailHandler returns a JSON post by id
func PostsDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["post_id"])
}

// PostsHandler returns an array of JSON posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	post := MockPost()
	fmt.Fprintf(w, "%s", post.ToJSON())
}

func main() {
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
