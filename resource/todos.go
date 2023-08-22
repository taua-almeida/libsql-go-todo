package resource

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/taua-almeida/libsql-go-todo/repository"
)

type TodosResourse struct{}

func (rs TodosResourse) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // GET /todos - read a list of todos
	r.Post("/", rs.Create) // POST /todos - create a todo

	return r
}

func (rs TodosResourse) List(w http.ResponseWriter, r *http.Request) {
	todos, err := repository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (rs TodosResourse) Create(w http.ResponseWriter, r *http.Request) {
	todo := repository.Todo{}
	json.NewDecoder(r.Body).Decode(&todo)

	errors, err := todo.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
