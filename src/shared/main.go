package shared

import (
	"darklab_training_postgres/utils"
	"darklab_training_postgres/utils/types"
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)

func FixtureConn(dbname types.Dbname, callback func(dbname types.Dbname, conn *sql.DB)) {

	connStr := fmt.Sprintf("postgres://postgres:postgres@localhost/%s?sslmode=disable", dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	callback(dbname, db)
	fmt.Println("The database is connected")
}

func FixtureDatabase(callback func(dbname types.Dbname)) {
	new_dbname := types.Dbname(utils.TokenWord(8))
	FixtureConn("postgres", func(dbpath types.Dbname, conn *sql.DB) {

		_, err := conn.Exec(fmt.Sprintf("CREATE DATABASE %s", new_dbname))
		if err != nil {
			panic(err)
		}
		fmt.Println("supposedly created new db")

		callback(new_dbname)

		drop_connections := fmt.Sprintf("SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '%s'", new_dbname)
		_, err = conn.Exec(drop_connections)
		_ = err
		_, err = conn.Exec(fmt.Sprintf("DROP DATABASE %s", new_dbname))
		_ = err
	})
}

func FixtureConnTestDB(callback func(dbname types.Dbname, conn *sql.DB)) {
	FixtureDatabase(func(dbname types.Dbname) {
		FixtureConn(dbname, func(dbname types.Dbname, conn *sql.DB) {
			callback(dbname, conn)
		})
	})
}
