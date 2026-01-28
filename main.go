package main

import (
	"gopost-api/config"
	"gopost-api/server"
	"net/http"
)

/*
func Hola(w http.ResponseWriter, r *http.Request) {
	//age := r.PathValue("age")
	//name := r.PathValue("name")

	age := r.URL.Query().Get("age")
	name := r.URL.Query().Get("name")

	datos := map[string]string{
		"message": "Hello, Moral!",
		"name":    name,
		"age":     age,
		"status":  "success",
		"code":    "200",
		"detail":  "la cantidad de pc no son con moral",
	}

	// Convertir el mapa a JSON

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("{\"message\": \"Hello, Moral!\"}"))
	err := json.NewEncoder(w).Encode(datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
*/

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\": \"healthy\"}"))
}

func main() {
	/*
		//http.HandleFunc("/moral", Hola)
		mux := http.NewServeMux()
		//mux.HandleFunc("GET /moral/info", Hola)
		mux.HandleFunc("GET /healt", health)
		mux.HandleFunc("GET /posts", handlers.GetPosts)
		mux.HandleFunc("POST /posts", handlers.CreatePost)
		mux.HandleFunc("GET /posts/{id}", handlers.GetPostByID)

		fmt.Println("Starting server on http://localhost:8080")
		http.ListenAndServe(":8080", mux)
	*/
	config := config.LoadConfig()
	app := server.NewApp()
	app.Runserver(config.Port)

}
