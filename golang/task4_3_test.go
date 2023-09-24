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

func TestTask4Query3(t *testing.T) {
	RunSubTests("4_3", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
			N := 50
			result := conn_orm.Raw(
				Task4Query3,
				sql.Named("N", N),
			)
			utils.Check(result.Error)

			count := CountRows(result)
			assert.Equal(t, N, count)
		})
	})
}
