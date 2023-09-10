package task3

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils/types"
	"database/sql"
	_ "embed"
	"testing"

	"darklab_training_postgres/src/task2"

	"gorm.io/gorm"
)

var (
	//go:embed queries/query3_1.sql
	Query1 string

	//go:embed queries/query3_2.sql
	Query2 string

	//go:embed queries/query3_3.sql
	Query3 string

	//go:embed queries/query3_4.sql
	Query4 string

	//go:embed queries/query3_5.sql
	Query5 string
)

func TestQuery1(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		task2.FixtureTask2Migrations(conn)

		task2.FixtureFillWithData(
			dbname,
			types.MaxUsers(1000),
		)
	})
}
