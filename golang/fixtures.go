package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Migration1          string
	Migration2          string
	MigrationAddIndexes string
)

func init() {
	Migration1 = utils.ReadProjectFile("sql/task2/migrations/task2_1.sql")
	Migration2 = utils.ReadProjectFile("sql/task2/migrations/task2_2_disable_triggers_for_tests.sql")
	MigrationAddIndexes = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/migrations/task3_7.sql"))
}

func FixtureTask2Migrations(conn *sql.DB) {
	utils.MustExec(conn, Migration1)
	// utils.MustExec(conn, Migration2)
}

func FixtureTask3Migrations(conn_orm *gorm.DB) {
	res := conn_orm.Raw(MigrationAddIndexes)
	if res.Error != nil {
		panic(res.Error)
	}
}

func FixtureFillWithData(
	dbname types.Dbname,
	max_users types.MaxUsers,
	posts_per_user types.PostsPerUser,
) {
	PostsAmount := int(max_users) * int(posts_per_user)
	PostVisitCount := 5000
	shared.FixtureTimeMeasure(func() {
		user_bulker := shared.Bulker[model.User]{
			Amount_to_create: types.AmountCreate(max_users),
			Bulk_max:         types.BulkMax(8000),
			Dbname:           dbname,
		}
		user_bulker.Init().BulkCreate(func(u *model.User) { u.Fill() })

		post_bulker := shared.Bulker[model.Post]{
			Amount_to_create: types.AmountCreate(PostsAmount),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           dbname,
		}

		userCounter := 1
		post_bulker.Init().BulkCreate(func(p *model.Post) {
			p.Fill(userCounter)
			userCounter++
			if userCounter > int(max_users) {
				userCounter = 1
			}
		})

		post_visit_bulker := shared.Bulker[model.PostVisits]{
			Amount_to_create: types.AmountCreate(PostVisitCount),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           dbname,
		}
		postCounter := 1
		post_visit_bulker.Init().BulkCreate(func(p *model.PostVisits) {
			p.Fill(postCounter)
			postCounter++
			if postCounter > int(max_users) {
				postCounter = 1
			}
		})

	}, "database filling with data")
}
