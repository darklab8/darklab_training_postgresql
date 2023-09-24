package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestSelect1(t *testing.T) {
	shared.FixtureConn(testdb.UnitTests.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		rows, err := conn.Query("SELECT 1")
		if err != nil {
			panic(err)
		}
		rows.Next()
		var answer int
		rows.Scan(&answer)
		fmt.Println(answer)

		assert.Equal(t, 1, answer)
	})
}
