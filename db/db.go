package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func Init(username, password, database, host, port string) (Database, error) {
	db := Database{}
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	connection, err := sql.Open("postgres", dataSource)
	if err != nil {
		return db, err
	}
	db.Connection = connection
	err = db.Connection.Ping()
	if err != nil {
		return db, err
	}

	initQuery := `CREATE TABLE IF NOT EXISTS urls
					(
						short_url TEXT PRIMARY KEY,
						url       TEXT
					);`
	_, err = db.Connection.Query(initQuery)
	if err != nil {
		return Database{}, err
	}
	return db, nil
}
