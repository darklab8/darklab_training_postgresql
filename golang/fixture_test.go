package golang

import (
	"darklab_training_postgres/golang/settings"
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

/*
This function signature must be located in file ending with _test.go :)
*/
func TestMain(m *testing.M) {
	fmt.Println("seting")
	var code int

	testdbname := types.Dbname(utils.TokenWord(8))
	shared.FixtureConnTestDBWithName(types.AutodestroyDB(true), types.Dbname(fmt.Sprintf("%s_unit_tests", testdbname)), func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		// Unit tests db
		FixtureTask2Migrations(conn)
		testdb.UnitTests.Dbname = dbname
		FixtureFillWithData(testdb.UnitTests)
		_, err := conn.Exec("REFRESH MATERIALIZED VIEW user_ratings")
		utils.Check(err)
		FixtureTask3Migrations(conn)

		shared.FixtureConnTestDBWithName(types.AutodestroyDB(false), types.Dbname(fmt.Sprintf("%s_indexes", testdbname)), func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

			if settings.ENABLED_PERFORMANCE_TESTS {
				FixtureTask2Migrations(conn)
				testdb.PerformanceWithIndexes.Dbname = dbname
				FixtureFillWithData(testdb.PerformanceWithIndexes)
				_, err := conn.Exec("REFRESH MATERIALIZED VIEW user_ratings")
				utils.Check(err)

				// Indexless test db. Not autodestroyed
				testdb.PerformanceIndexless.Dbname = types.Dbname(fmt.Sprintf("%s_indexless", testdbname))
				_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE %s", testdb.PerformanceIndexless.Dbname, dbname))
				utils.Check(err)

				FixtureTask3Migrations(conn)

			}
			code = m.Run()
		})
	})
	fmt.Println("teardown")

	os.Exit(code)
}
