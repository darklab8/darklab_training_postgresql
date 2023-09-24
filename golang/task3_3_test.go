package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestTask3Query3PostsAwaitingPublishing(t *testing.T) {
	RunSubTests("3_3", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

			N := (rand.Intn(int(db_params.MaxUsers)+int(db_params.PostsPerUser)) / 4)
			result := conn_orm.Raw(Task3Query3, sql.Named("N", N))
			if result.Error != nil {
				panic(result.Error)
			}

			rows, err := result.Rows()
			if err != nil {
				panic(err)
			}

			count_rows := 0
			for rows.Next() {
				count_rows += 1
			}
			assert.Equal(t, N, count_rows)
		})
	})
}
