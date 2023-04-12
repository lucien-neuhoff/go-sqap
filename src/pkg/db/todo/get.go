package db_todo

import (
	"errors"
	db_user "go-sql/pkg/db/user"
	helper "go-sql/pkg/helper"
	"log"
)

func Get() []helper.Todo {
	var todos []helper.Todo

	rows, err := helper.DB.Query("SELECT * FROM todos")
	if err != nil {
		log.Printf("db.Get: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo helper.Todo
		if err := rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Complete); err != nil {
			log.Printf("db.Get: %s", err)
		}
		todos = append(todos, todo)
	}

	return todos
}

func GetById(todo_id string, user_id string, session_key string) (helper.Todo, error) {
	var todo = helper.Todo{ID: todo_id}

	valid, err := db_user.VerifySessionKey(user_id, session_key)
	if err != nil {
		log.Println("db.todo.GetById::VerifySessionKey: ", err)
		return todo, err
	}
	if !valid {
		return todo, errors.New("unvalid session")
	}

	err = helper.DB.QueryRow("SELECT * FROM todos WHERE id=?", todo_id).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Complete)
	if err != nil {
		return todo, err
	}

	return todo, nil
}
