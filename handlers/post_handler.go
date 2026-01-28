package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// handlers/post_handler.go creando la estructura del json
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var posts = []Post{}
var idCounter = 1

// GetPosts maneja la obtención de todos los posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

// CreatePost maneja la creación de un nuevo post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// asignar los nuevos valores al post
	newPost.ID = idCounter
	idCounter++
	posts = append(posts, newPost)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newPost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for _, post := range posts {
		if post.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}
