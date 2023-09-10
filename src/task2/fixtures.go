package task2

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils"
	"darklab_training_postgres/utils/types"
	"database/sql"
	_ "embed"
	"fmt"
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

func FixtureFillWithData(dbname types.Dbname, max_users types.MaxUsers) {
	shared.FixtureTimeMeasure(func() {
		bulker := shared.Bulker[User]{
			Amount_to_create: types.AmountCreate(max_users),
			Bulk_max:         types.BulkMax(8000),
			Dbname:           dbname,
		}
		bulker.Init().BulkCreate(func(u *User) { u.Fill() })
	}, "database filling with data")
}
