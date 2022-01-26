package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.HandleFunc("/", helloWeb)

	PORT := os.Getenv("HTTP_PORT")
	if PORT == "" {
		PORT = "3000"
	}

	http.ListenAndServe(":"+PORT, r)
}

func helloWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Edu Mano!"))
}
