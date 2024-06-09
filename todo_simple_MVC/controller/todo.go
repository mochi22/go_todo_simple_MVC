// controller/todo.go
package controller

import (
    "html/template"
    "net/http"
    "strconv"
    "todo/model"
)

type ViewData struct {
    Todos []model.Todo
}

func prepareData(todos []model.Todo) ViewData {
    // ViewDataという構造体を作成し、todosスライスを格納する
    data := ViewData{
        Todos: todos,
    }
    return data
}

var todos []model.Todo

var Current_id int

func renderTemplate(w http.ResponseWriter, tmpl string, data ViewData) error {
    t, err := template.ParseFiles("view/" + tmpl)
    if err != nil {
        return err
    }
    err = t.Execute(w, data)
    if err != nil {
        return err
    }
    return nil
}

// Todo項目の追加
func AddTodo(w http.ResponseWriter, r *http.Request) {
    // フォームから送信されたデータを取得する
    item := getItem(r)
	// 新しいTodoを作成し、モデルに追加
	todo := model.Todo{
        ID: Current_id,
		Item: item,
	}
	todos = append(todos, todo)
    Current_id += 1
	// 一覧ページにリダイレクトする
	http.Redirect(w, r, "/todo", http.StatusFound)
}

// Todo一覧の取得
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := todos
    data := prepareData(todos)
    err := renderTemplate(w, "todos.html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Todo項目の更新
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    // フォームから送信されたデータ(ID、Item、Completed)を取得する
    id := getId(r)
    item := getItem(r)
    completed := getCompleted(r)

    // モデルのTodoを更新する
    for i, t := range todos {
        if t.ID == id {
            todos[i].Item = item
            todos[i].Completed = completed
            break
        }
    }

    // 一覧ページにリダイレクトする
    http.Redirect(w, r, "/todo", http.StatusFound)
}

// Todo項目の削除
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
        // フォームから送信されたデータ(ID)を取得する
        id := getId(r)

		// モデルからTodoを削除する
		var newTodos []model.Todo
		for _, t := range todos {
			if t.ID != id {
				newTodos = append(newTodos, t)
			} else{
                println("delete", t.ID, t.Item)
            }
		}
		todos = newTodos
	
		// 一覧ページにリダイレクトする
		http.Redirect(w, r, "/todo", http.StatusFound)
}


func getId(r *http.Request) int {
    id, _ := strconv.Atoi(r.FormValue("id"))
    return id
}

func getItem(r *http.Request) string {
    return r.FormValue("item")
}

func getCompleted(r *http.Request) bool {
    completed := r.FormValue("completed") == "true"
    return completed
}



