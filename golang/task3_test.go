package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

var (
	Task3Query1 string
	Task3Query2 string
	Task3Query3 string
	Task3Query4 string
	Task3Query5 string
)

func init() {
	Task3Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_1.sql"))
	Task3Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_2.sql"))
	Task3Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_3.sql"))
	Task3Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_4.sql"))
	Task3Query5 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_5.sql"))
}

func TestQueryReuseSetup2(t *testing.T) {
	shared.FixtureConn(testdb.UnitTests.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		var count int64
		conn_orm.Model(&model.Post{}).Count(&count)
		assert.Equal(t, int(testdb.UnitTests.MaxUsers)*int(testdb.UnitTests.PostsPerUser), int(count))
	})
}

func TestMigration(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		FixtureTask2Migrations(conn)
		FixtureTask3Migrations(conn)
	})
}
