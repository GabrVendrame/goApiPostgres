package db

import (
	"database/sql"
	"fmt"
	"postgresApi/configs"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error){
	conf := configs.GetDb()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

    connection, err := sql.Open("postgres", sc)

    if err != nil {
        panic(err)
    }

    err = connection.Ping()

    return connection, err
}
