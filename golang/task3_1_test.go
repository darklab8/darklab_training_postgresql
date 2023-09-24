package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestTask3Query1UserPostCount(t *testing.T) {
	RunSubTests("3_1", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
			author_id := 1 + rand.Intn(int(db_params.MaxUsers)-1)
			result := conn_orm.Raw(Task3Query1, sql.Named("author_id", author_id))
			utils.Check(result.Error)

			var count int
			result.Scan(&count)
			assert.Equal(t, int(db_params.PostsPerUser), count)
		})
	})
}
