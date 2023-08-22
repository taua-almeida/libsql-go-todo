package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/taua-almeida/libsql-go-todo/database"
	"github.com/taua-almeida/libsql-go-todo/resource"
)

func main() {
	err := database.InitDB("file:////Users/tauaalmeida/Documents/go-lang/libsql-go-todo/todo.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is up"))
	})

	r.Mount("/todos", resource.TodosResourse{}.Routes())

	http.ListenAndServe(":42069", r)
}
