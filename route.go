package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID    int    `json: "id"`
	Title string `json: "title"`
	Text  string `json: "text"`
}

var (
	posts []*Post
)

func init() {
	posts = append(posts, 
		&Post{ID: 1, Title: "Title 1", Text: "Text 1"},
		&Post{ID: 2, Title: "Title 2", Text: "Text 2"},
		&Post{ID: 3, Title: "Title 3", Text: "Text 3"},
	)
}

func listPosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurs when marshaling posts"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func addPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{}
	err := json.NewDecoder(req.Body).Decode(post)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occurs when unmarshaling post from request"))
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)

	w.WriteHeader(http.StatusCreated)
	result, err := json.Marshal(post)
	w.Write(result)
}
