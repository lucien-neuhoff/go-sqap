package db_todo

import (
	"fmt"
	"go-sql/pkg/helper"
	"log"
)

func Post(todo *helper.Todo) {
	_, err := helper.DB.Query("INSERT INTO todos (id, title, description, complete) VALUES ('" + todo.ID + "','" + todo.Title + "','" + todo.Description + "','" + fmt.Sprint(todo.Complete) + "')")
	if err != nil {
		log.Println(err)
	}
}
