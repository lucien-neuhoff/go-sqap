package db_todo

import (
	"go-sql/pkg/helper"
	"log"
)

func Post(todo *helper.Todo) {
	log.Printf("db.todo.Post: posting todo with id '%s'", todo.ID)
	_, err := helper.DB.Query("INSERT INTO todos (id, title, description, complete, user_id) VALUES (?, ?, ?, ?, ?)", todo.ID, todo.Title, todo.Description, todo.Complete, todo.UserID)
	if err != nil {
		log.Println(err)
	}
}
