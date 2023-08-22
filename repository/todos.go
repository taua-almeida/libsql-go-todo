package repository

import "github.com/taua-almeida/libsql-go-todo/database"

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Id          int    `json:"id"`
}

type ValidateErrorMap = map[string]string

func (t *Todo) validate() ValidateErrorMap {
	errors := ValidateErrorMap{}

	if t.Title == "" {
		errors["title"] = "Title is required"
	}

	return errors
}

func (t *Todo) Create() (ValidateErrorMap, error) {
	errors := t.validate()

	if len(errors) > 0 {
		return errors, nil
	}

	_, err := database.Db.Exec("INSERT INTO todo (title, description, completed) VALUES (?, ?, ?)", t.Title, t.Description, t.Completed)

	return errors, err
}

func (t *Todo) Update() (ValidateErrorMap, error) {
	errors := t.validate()

	if len(errors) > 0 {
		return errors, nil
	}

	_, err := database.Db.Exec("UPDATE todo SET title = ?, description = ?, completed = ? WHERE id = ?", t.Title, t.Description, t.Completed, t.Id)

	return errors, err
}

func (t *Todo) Delete() error {
	_, err := database.Db.Exec("DELETE FROM todo WHERE id = ?", t.Id)

	return err
}

func FindAll() ([]Todo, error) {
	rows, err := database.Db.Query("SELECT id, title, description, completed FROM todo")

	if err != nil {
		return nil, err
	}

	var todos []Todo

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
		todos = append(todos, todo)
	}

	return todos, nil
}
