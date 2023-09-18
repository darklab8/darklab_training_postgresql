package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type TemporalDB struct {
	MaxUsers     types.MaxUsers
	PostsPerUser types.PostsPerUser
	Dbname       types.Dbname
}

var (
	TempDb = TemporalDB{
		MaxUsers:     10000,
		PostsPerUser: 50,
	}
)

func FixtureFillTemporalDB(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
	FixtureTask2Migrations(conn)
	FixtureFillWithData(
		dbname,
		TempDb.MaxUsers,
		TempDb.PostsPerUser,
	)
	TempDb.Dbname = dbname
	FixtureTask3Migrations(conn_orm, bundb)
}

func TestMain(m *testing.M) {
	fmt.Println("seting")
	var code int
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		FixtureFillTemporalDB(dbname, conn, conn_orm, bundb)
		code = m.Run()
		fmt.Println("teardown")
	})
	os.Exit(code)
}
