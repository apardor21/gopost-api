package handlers

import (
	"gopost-api/server"
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
func GetPosts(c *server.Context) {
	error := c.JSON(http.StatusOK, posts)
	if error != nil {
		http.Error(c.Rwrite, error.Error(), http.StatusInternalServerError)
	}
}

// CreatePost maneja la creación de un nuevo post
func CreatePost(c *server.Context) {

	var newPost Post
	err := c.BindJSON(newPost)
	if err != nil {
		http.Error(c.Rwrite, "Error al decodificar Json", http.StatusBadRequest)
		return
	}
	// asignar los nuevos valores al post
	newPost.ID = idCounter
	idCounter++
	posts = append(posts, newPost)

	err = c.JSON(http.StatusCreated, newPost)

	if err != nil {
		http.Error(c.Rwrite, "Error al codificar Json", http.StatusBadRequest)
		return
	}

}

func GetPostByID(c *server.Context) {

	idStr := c.Request.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	for _, post := range posts {
		if post.ID == id {
			c.JSON(http.StatusOK, post)
			return
		}
	}
	http.Error(c.Rwrite, "Post not found", http.StatusNotFound)
}
