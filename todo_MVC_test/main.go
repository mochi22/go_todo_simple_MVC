// main.go
package main

import (
	"net/http"

	"github.com/mochi22/go_tools/tree/main/todo_MVC_test/controller"
)

func main() {
	mispell := "misspell"
	println(mispell)

	http.HandleFunc("/todo", controller.GetTodos)
	http.HandleFunc("/todo/add", controller.AddTodo)
	http.HandleFunc("/todo/update", controller.UpdateTodo)
	http.HandleFunc("/todo/delete", controller.DeleteTodo)

	http.ListenAndServe(":8080", nil)
}
