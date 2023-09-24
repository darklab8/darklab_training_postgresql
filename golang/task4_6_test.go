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

func TestTask4Query6(t *testing.T) {
	RunSubTests("4_6", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
			N := 50
			result := conn_orm.Raw(
				Task4Query6,
				sql.Named("N", N),
			)
			utils.Check(result.Error)

			tag_count := 3
			count := CountRows(result)
			assert.Equal(t, tag_count, count)
		})
	})
}
