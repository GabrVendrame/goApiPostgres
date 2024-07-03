package models

import "postgresApi/db"

func Get(id int64) (todo Todo, err error) {
	connection, err := db.Connect()

	if err != nil {
		return
	}
	defer connection.Close()

    sql := `SELECT * FROM todos WHERE id=$1`

	row := connection.QueryRow(sql, id)

	err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

	return
}

func GetAll() (todos []Todo, err error) {
	connection, err := db.Connect()
	if err != nil {
		return
	}
	defer connection.Close()

    sql := `SELECT * FROM todos`

    rows, err := connection.Query(sql)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
