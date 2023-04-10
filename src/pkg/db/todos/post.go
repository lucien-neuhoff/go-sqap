package db_todos

import (
	"fmt"
	"go-sql/pkg/helper"
	"log"
)

func Post(todo *helper.Todo) {
	log.Println(todo.ID)
	_, err := helper.DB.Query("INSERT INTO todos (id, title, description, complete) VALUES ('" + todo.ID + "','" + todo.Title + "','" + todo.Description + "','" + fmt.Sprint(todo.Complete) + "')")
	if err != nil {
		log.Fatal(err)
	}
}
