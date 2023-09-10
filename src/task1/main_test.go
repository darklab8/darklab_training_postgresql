package task1

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils/types"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	shared.FixtureConnTestDB(func(dbpath types.Dbname, conn *sql.DB) {
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
	})
}
