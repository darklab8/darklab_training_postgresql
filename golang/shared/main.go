package shared

import (
	"crypto/tls"
	"darklab_training_postgres/golang/settings"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"log"
	"strings"
	"time"

	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func FixtureConn(dbname types.Dbname, callback func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB)) {

	connStr := fmt.Sprintf("postgres://postgres:postgres@%s/%s?sslmode=disable", settings.DatabaseHost, dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	defer db.Close()

	timeout := time.Duration(30) * time.Second
	bundb1 := sql.OpenDB(pgdriver.NewConnector(
		// pgdriver.WithDSN(connStr)
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf(settings.DatabaseHost+":5432")),
		pgdriver.WithUser("postgres"),
		pgdriver.WithPassword("postgres"),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithInsecure(true),
		pgdriver.WithDatabase(string(dbname)),
		pgdriver.WithTimeout(timeout),
		pgdriver.WithDialTimeout(timeout),
		pgdriver.WithReadTimeout(timeout),
		pgdriver.WithWriteTimeout(timeout),
	))
	bundb := bun.NewDB(bundb1, pgdialect.New())
	defer bundb.Close()

	callback(dbname, db, gormDB, bundb)
}

func FixtureDatabase(callback func(dbname types.Dbname)) {
	new_dbname := types.Dbname(utils.TokenWord(8))
	FixtureConn("postgres", func(dbpath types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

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

func FixtureConnTestDB(callback func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB)) {
	FixtureDatabase(func(dbname types.Dbname) {
		FixtureConn(dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
			callback(dbname, conn, conn_orm, bundb)
		})
	})
}

func FixtureTimeMeasure(callback func(), msgs ...string) {
	start := time.Now()

	callback()

	elapsed := time.Since(start)
	log.Printf("time elapsed %s for %s", elapsed, strings.Join(msgs, " "))
}
