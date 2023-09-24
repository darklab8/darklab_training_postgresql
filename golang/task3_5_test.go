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

func TestTask3Query5FindNPostsWithMostRating(t *testing.T) {
	RunSubTests("3_5", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

			result := conn_orm.Raw(Task3Query5,
				sql.Named("N", 20),
			)

			utils.Check(result.Error)
			count_rows := CountRows(result)

			assert.Equal(t, 20, count_rows)
		})
	})
}
