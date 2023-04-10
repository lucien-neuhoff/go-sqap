package helper

import "database/sql"

var DB *sql.DB
var ENVS map[string]string

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    int    `json:"complete"`
}

type Message struct {
	From    string `json:"from"`
	Content string `json:"message"`
}
