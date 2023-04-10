package db_todo

import (
	"fmt"
	helper "go-sql/pkg/helper"
	"log"
)

func Get() []helper.Todo {
	var todos []helper.Todo

	rows, err := helper.DB.Query("SELECT * FROM todos")
	if err != nil {
		log.Println("db.Get: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo helper.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Complete); err != nil {
			log.Println("db.Get: ", err)
		}
		todos = append(todos, todo)
	}

	return todos
}

func GetById(id string) (helper.Todo, string) {
	var todo helper.Todo

	err := helper.DB.QueryRow("SELECT * FROM todos WHERE id="+id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Complete)
	if err != nil {
		log.Println("db.todo.GetById: ", err)
		return todo, fmt.Sprintf("Todo with id '%s' not found", id)
	}

	return todo, ""
}
