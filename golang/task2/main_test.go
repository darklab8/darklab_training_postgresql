package task2

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateData(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		FixtureTask2Migrations(conn)

		max_users := 10000
		FixtureFillWithData(
			dbname,
			types.MaxUsers(max_users),
			types.PostsPerUser(50),
		)

		var count int
		rows, _ := conn.Query("SELECT count(*) FROM user_")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, count, max_users)

		rows, _ = conn.Query("SELECT count(*) FROM post")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, count, max_users*50)

	})
}
