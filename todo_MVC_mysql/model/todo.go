package model

type Todo struct {
	ID        int
	Item      string
	Completed bool
}

// var todos []Todo

func AddTodoDB(id int, item string, completed bool) error {
	_, err := DB.Exec("INSERT INTO todo (id, item, completed) VALUES (?, ?, ?)", id, item, completed)
	if err != nil {
		return err
	}
	return nil
}

func GetTodosDB() ([]Todo, error) {
	rows, err := DB.Query("SELECT * FROM todo")
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
	_, err := DB.Exec("UPDATE todo SET item=?, completed=? WHERE id=?", item, completed, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodoDB(id int) error {
	_, err := DB.Exec("DELETE FROM todo WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
