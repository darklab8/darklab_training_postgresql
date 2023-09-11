package utils

import (
	"database/sql"
	"strings"
)

func MustExec(conn *sql.DB, query string) {
	_, err := conn.Exec(query)
	if err != nil {
		panic(err)
	}
}

func GetSQLFile(data string) string {
	return strings.ReplaceAll(data, ":", "@")
}
