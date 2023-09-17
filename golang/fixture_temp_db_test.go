package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	"fmt"
	"os"
	"testing"

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

func FixtureFillTemporalDB(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
	FixtureTask2Migrations(conn)
	FixtureFillWithData(
		dbname,
		TempDb.MaxUsers,
		TempDb.PostsPerUser,
	)
	TempDb.Dbname = dbname

	res := conn_orm.Raw(MigrationAddIndexes)
	if res.Error != nil {
		panic(res.Error)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("seting")
	var code int
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		FixtureFillTemporalDB(dbname, conn, conn_orm)
		code = m.Run()
		fmt.Println("teardown")
	})
	os.Exit(code)
}
