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

func TestTask3Query4RecentlyUpdatedPostsByTag(t *testing.T) {
	RunSubTests("3_4", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

			raw_sql := conn_orm.ToSQL(func(tx *gorm.DB) *gorm.DB {
				return tx.Raw(Task3Query4,
					sql.Named("N", 20),
					sql.Named("K", 1),
					sql.Named("L", 10),
					sql.Named("tag", "a1"),
				)
			})
			result, err := conn.Exec(raw_sql)

			utils.Check(err)

			count_rows, err := result.RowsAffected()
			utils.Check(err)

			assert.GreaterOrEqual(t, int(10), int(count_rows))
		})
	})
}
