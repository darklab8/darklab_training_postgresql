package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
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
	TempDbIndexless = TemporalDB{
		MaxUsers:     10000,
		PostsPerUser: 50,
	}
)

func FixtureFillTemporalDB(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

}

func TestMain(m *testing.M) {
	fmt.Println("seting")
	var code int
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		FixtureTask2Migrations(conn)
		FixtureFillWithData(
			dbname,
			TempDb.MaxUsers,
			TempDb.PostsPerUser,
		)
		_, err := conn.Exec("REFRESH MATERIALIZED VIEW user_ratings")
		utils.Check(err)
		TempDb.Dbname = dbname

		TempDbIndexless.Dbname = types.Dbname(utils.TokenWord(8))

		// _, err = conn.Exec(fmt.Sprintf("SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '%s'", dbname))
		// utils.Check(err)
		_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE %s", TempDbIndexless.Dbname, dbname))
		utils.Check(err)

		FixtureTask3Migrations(conn_orm, bundb)

		code = m.Run()
		fmt.Println("teardown")
	})
	os.Exit(code)
}
