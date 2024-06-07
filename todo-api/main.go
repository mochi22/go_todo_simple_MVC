package main

import (
    "log"
    "net/http"
    "todo-api/src/handlers"
)

func main() {
	http.HandleFunc("/todos", handlers.TodoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}