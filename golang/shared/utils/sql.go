package utils

import (
	"database/sql"
)

func MustExec(conn *sql.DB, query string) {
	_, err := conn.Exec(query)
	if err != nil {
		panic(err)
	}
}

func GetSQLFile(data string) string {
	targetSymbol := rune(':')
	runes := []rune(data)
	for pos, char := range runes {

		if char != targetSymbol {
			continue
		}

		if runes[pos-1] == targetSymbol || runes[pos+1] == targetSymbol {
			continue
		}

		runes[pos] = rune('@')
	}
	return string(runes)
}
