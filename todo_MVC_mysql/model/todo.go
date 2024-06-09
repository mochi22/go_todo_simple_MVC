package model

type Todo struct {
	ID        int
	Item      string
	Completed bool
}

// var todos []Todo

func AddTodoDB(item string, completed bool) error {
	_, err := DB.Exec("INSERT INTO todos (item, completed) VALUES (?, ?)", item, completed)
	if err != nil {
		return err
	}
	return nil
}

func GetTodosDB() ([]Todo, error) {
	rows, err := DB.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Item, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func UpdateTodoDB(id int, item string, completed bool) error {
	_, err := DB.Exec("UPDATE todos SET item=?, completed=? WHERE id=?", item, completed, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodoDB(id int) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
