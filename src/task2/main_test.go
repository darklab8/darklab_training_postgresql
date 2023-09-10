package task2

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils"
	"darklab_training_postgres/utils/types"
	"database/sql"
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	//go:embed migrations/task2_1.sql
	Migration1 string

	//go:embed migrations/task2_2_disable_triggers_for_tests.sql
	Migration2 string
)

func FixtureTask2Migrations(conn *sql.DB) {
	fmt.Println()

	utils.MustExec(conn, Migration1)
	utils.MustExec(conn, Migration2)
}

func TestCreateData(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		FixtureTask2Migrations(conn)

		max_users := 1000000

		shared.FixtureTimeMeasure(func() {
			for i := 0; i < 10; i++ {
				shared.BulkCreate[User](dbname, types.AmountCreate(max_users/10), types.BulkMax(8000), conn_orm, func(u *User) {
					u.Fill()
				})
			}

			var user_count int
			rows, _ := conn.Query("SELECT count(*) FROM user_")
			rows.Next()
			rows.Scan(&user_count)
			assert.Equal(t, user_count, max_users)
		}, "database filling with data")

	})
}
