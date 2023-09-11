package task1

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
	shared.FixtureConnTestDB(func(dbpath types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		fmt.Println(dbpath)
		rows, err := conn.Query("SELECT 1")
		if err != nil {
			panic(err)
		}
		rows.Next()
		var answer int
		rows.Scan(&answer)
		fmt.Println(answer)

		assert.Equal(t, 1, answer)
		fmt.Println("running testmigrate")
	})
}