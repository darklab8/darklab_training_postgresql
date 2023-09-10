package task2

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils/types"
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
		)

		var user_count int
		rows, _ := conn.Query("SELECT count(*) FROM user_")
		rows.Next()
		rows.Scan(&user_count)
		assert.Equal(t, user_count, max_users)

	})
}
