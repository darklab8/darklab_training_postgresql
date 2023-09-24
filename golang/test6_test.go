package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"testing"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

var (
	Task6Query1 string
	Task6Query2 string
	Task6Query3 string
	Task6Query4 string
)

func init() {
	Task6Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task6/queries/query_6_1.sql"))
	Task6Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task6/queries/query_6_2.sql"))
	Task6Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task6/queries/query_6_3.sql"))
	Task6Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task6/queries/query_6_4.sql"))
}

func TestTask6Query1(t *testing.T) {
	shared.FixtureConn(testdb.UnitTests.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		result := conn_orm.Raw(
			Task6Query1,
			sql.Named("post_id", 10),
		)
		utils.Check(result.Error)
	})
}

func TestTask6Query2(t *testing.T) {
	shared.FixtureConn(testdb.UnitTests.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		result := conn_orm.Raw(
			Task6Query2,
			sql.Named("N", 10),
		)
		utils.Check(result.Error)
	})
}

func TestTask6Query3(t *testing.T) {
	shared.FixtureConnTestDB(types.AutodestroyDB(true), func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		FixtureTask2Migrations(conn)
		FixtureTask3Migrations(conn)

		result := conn_orm.Raw(Task6Query3)
		utils.Check(result.Error)
	})
}

func TestTask6Query4(t *testing.T) {
	shared.FixtureConnTestDB(types.AutodestroyDB(true), func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		res := conn_orm.Raw(`
		CREATE TABLE post
		(
			id SERIAL PRIMARY KEY,
		)
		`)
		utils.Check(res.Error)
		result := conn_orm.Raw(Task6Query4)
		utils.Check(result.Error)
	})
}
