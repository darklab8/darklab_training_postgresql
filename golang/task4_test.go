package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	Task4Query1   string
	Task4Query1_2 string
	Task4Query2   string
	Task4Query3   string
	Task4Query4   string
	Task4Query5   string
	Task4Query6   string
)

func init() {
	Task4Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_1.sql"))
	Task4Query1_2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_1_2.sql"))
	Task4Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_2.sql"))
	Task4Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_3.sql"))
	Task4Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_4.sql"))
	Task4Query5 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_5.sql"))
	Task4Query6 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_6.sql"))
}

func CountRows(result *gorm.DB) int {
	rows, err := result.Rows()
	if err != nil {
		panic(err)
	}
	count_rows := 0
	for rows.Next() {
		count_rows += 1
	}
	return count_rows
}

func TestTask4Query1(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		Query1Test := func(query1 string) {
			N := 50
			result := conn_orm.Raw(query1, sql.Named("N", N))
			utils.Check(result.Error)

			count := CountRows(result)
			assert.Equal(t, int(TempDb.PostsPerUser), count)
		}

		t.Run("taskquery1_1", func(t *testing.T) {
			Query1Test(Task4Query1)
		})
		t.Run("taskquery1_2", func(t *testing.T) {
			Query1Test(Task4Query1_2)
		})
	})
}
