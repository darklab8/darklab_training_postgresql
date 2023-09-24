package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestTask4Query1MostVisitedPostInAYear(t *testing.T) {
	RunSubTests("4_1", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

			Query1Test := func(query1 string) {
				N := 50
				result := conn_orm.Raw(query1, sql.Named("N", N))
				utils.Check(result.Error)

				count := CountRows(result)
				assert.Equal(t, int(db_params.PostsPerUser), count)
			}

			t.Run("4_1_1", func(t *testing.T) {
				Query1Test(Task4Query1)
			})
			t.Run("4_1_2", func(t *testing.T) {
				Query1Test(Task4Query1_2)
			})
		})
	})
}
