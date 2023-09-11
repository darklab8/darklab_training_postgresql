package utils

import "database/sql"

func MustExec(conn *sql.DB, query string) {
	_, err := conn.Exec(query)
	if err != nil {
		panic(err)
	}
}
