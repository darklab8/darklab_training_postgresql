package section1

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils/types"
	"database/sql"
	"fmt"
	"testing"
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
	})
}
