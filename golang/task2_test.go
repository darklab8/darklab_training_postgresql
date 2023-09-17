package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateData(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		var count int
		rows, _ := conn.Query("SELECT count(*) FROM user_")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, int(TempDb.MaxUsers), count)

		rows, _ = conn.Query("SELECT count(*) FROM post")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, int(TempDb.MaxUsers)*int(TempDb.PostsPerUser), count)
	})
}
