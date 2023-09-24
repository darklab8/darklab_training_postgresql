package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"strings"
)

var (
	Migration1          string
	Migration2          string
	MigrationAddIndexes string
)

func init() {
	Migration1 = utils.ReadProjectFile("sql/task2/migrations/task2_1.sql")
	MigrationAddIndexes = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/migrations/task3_7.sql"))
}

func FixtureTask2Migrations(conn *sql.DB) {
	utils.MustExec(conn, Migration1)
	utils.MustExec(conn, Migration2)
}

func FixtureTask3Migrations(conn *sql.DB) {

	lines := strings.Split(MigrationAddIndexes, "\n")
	for _, line := range lines {

		if strings.HasPrefix(line, "--") {
			continue
		}
		if line == "" {
			continue
		}

		utils.MustExec(conn, line)
	}
}

type ModelCounter struct {
	count int
	max   int
}

func NewCounter(max int) ModelCounter {
	m := ModelCounter{
		count: 1,
		max:   max,
	}
	return m
}
func (m *ModelCounter) Next() int {
	defer func() {
		m.count += 1
		if m.count > int(m.max) {
			m.count = 1
		}
	}()
	return m.count
}

func FixtureFillWithData(
	params testdb.DBParams,
) {
	post_amount := int(params.MaxUsers) * int(params.PostsPerUser)

	shared.FixtureTimeMeasure(func() {
		user_bulker := shared.Bulker[model.User]{
			Amount_to_create: types.AmountCreate(params.MaxUsers),
			Bulk_max:         types.BulkMax(8000),
			Dbname:           params.Dbname,
		}
		user_bulker.Init().BulkCreate(func(u *model.User) { u.Fill() })

		post_bulker := shared.Bulker[model.Post]{
			Amount_to_create: types.AmountCreate(post_amount),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           params.Dbname,
		}
		user_counter := NewCounter(int(params.MaxUsers))
		post_bulker.Init().BulkCreate(func(p *model.Post) {
			p.Fill(user_counter.Next())
		})

		post_visit_bulker := shared.Bulker[model.PostVisits]{
			Amount_to_create: types.AmountCreate(params.PostVisits),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           params.Dbname,
		}
		post_counter := NewCounter(post_amount)
		post_visit_bulker.Init().BulkCreate(func(p *model.PostVisits) {
			p.Fill(post_counter.Next())
		})

		// TODO try to achieve high performance despite trigger for post_edition :/
		post_edition_bulker := shared.Bulker[model.PostEdition]{
			Amount_to_create: types.AmountCreate(params.PostEditions),
			Bulk_max:         types.BulkMax(4000),
			Dbname:           params.Dbname,
		}
		user_counter = NewCounter(int(params.MaxUsers))
		post_counter = NewCounter(post_amount)
		post_edition_bulker.Init().BulkCreate(func(p *model.PostEdition) {
			p.Fill(post_counter.Next(), user_counter.Next())
		})

	}, "database filling with data")
}
