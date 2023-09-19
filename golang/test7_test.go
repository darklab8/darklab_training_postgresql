package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"testing"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

var (
	Task7Query1 string
	Task7Query2 string
	Task7Query3 string
	Task7Query4 string
)

func init() {
	Task7Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_1.sql"))
	Task7Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_2.sql"))
	Task7Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_3.sql"))
	Task7Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_4.sql"))
}

func TestTask7Query1(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		result := conn_orm.Raw(Task7Query1)
		utils.Check(result.Error)
	})
}

func TestTask7Query2(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		result := conn_orm.Raw(Task7Query2)
		utils.Check(result.Error)
	})
}

func TestTask7Query3(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		result := conn_orm.Raw(Task7Query3)
		utils.Check(result.Error)
	})
}
