package store

import (
	"fmt"
	"time"
)

type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    DueDate   time.Time `json:"due_date"`
}

var currentID = 0
var todos = []Todo{}

func GetAllTodos() []Todo {
    return todos
}

func AddTodo(todo Todo) Todo {
    currentID++
    todo.ID = currentID
    todos = append(todos, todo)
    return todo
}

func RemoveTodo(id int) error {
    for i, todo := range todos {
		println("aaa:",i, todo.ID, todo.Title)
        if todo.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Todo with ID %d not found", id)
}