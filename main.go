// main.go
package main

import (
    "net/http"
    "todo/controller"
)

func main() {
    http.HandleFunc("/todo", controller.GetTodos)
    http.HandleFunc("/todo/add", controller.AddTodo)
    http.HandleFunc("/todo/update", controller.UpdateTodo)
    http.HandleFunc("/todo/delete", controller.DeleteTodo)

    http.ListenAndServe(":8080", nil)
}