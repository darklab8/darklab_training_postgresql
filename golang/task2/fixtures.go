package task2

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/orm"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"fmt"
)

var (
	Migration1 string
	Migration2 string
)

func init() {
	Migration1 = utils.ReadProjectFile("sql/task2/migrations/task2_1.sql")
	Migration2 = utils.ReadProjectFile("sql/task2/migrations/task2_2_disable_triggers_for_tests.sql")
}

func FixtureTask2Migrations(conn *sql.DB) {
	fmt.Println()

	utils.MustExec(conn, Migration1)
	utils.MustExec(conn, Migration2)
}

func FixtureFillWithData(
	dbname types.Dbname,
	max_users types.MaxUsers,
	posts_per_user types.PostsPerUser,
) {
	shared.FixtureTimeMeasure(func() {
		user_bulker := shared.Bulker[orm.User]{
			Amount_to_create: types.AmountCreate(max_users),
			Bulk_max:         types.BulkMax(8000),
			Dbname:           dbname,
		}
		user_bulker.Init().BulkCreate(func(u *orm.User) { u.Fill() })

		post_bulker := shared.Bulker[orm.Post]{
			Amount_to_create: types.AmountCreate(int(max_users) * int(posts_per_user)),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           dbname,
		}

		postCounter := 1
		post_bulker.Init().BulkCreate(func(p *orm.Post) {
			p.Fill(postCounter)
			postCounter++
			if postCounter > int(max_users) {
				postCounter = 1
			}
		})

	}, "database filling with data")
}
