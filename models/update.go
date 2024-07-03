package models

import "postgresApi/db"

func Update(id int64, todo Todo) (int64, error) {
	connection, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	sql := `UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`

	res, err := connection.Exec(sql, id, todo.Title, todo.Description, todo.Done)
    if err != nil {
        return 0, err
    }

	return res.RowsAffected()
}
