// main.go
package main

import (
	"log"
	"net/http"

	"todo/controller"
	"todo/model"
)

func main() {

	// データベースとの接続を確立する
	err := model.InitDB("username:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/todo", controller.GetTodos)
	http.HandleFunc("/todo/add", controller.AddTodo)
	http.HandleFunc("/todo/update", controller.UpdateTodo)
	http.HandleFunc("/todo/delete", controller.DeleteTodo)

	http.ListenAndServe(":8080", nil)
}
