package golang

import (
	"darklab_training_postgres/golang/settings"
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"fmt"
	"strings"
	"testing"
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
	dbname types.Dbname,
	max_users types.MaxUsers,
	posts_per_user types.PostsPerUser,
) {
	post_amount := int(max_users) * int(posts_per_user)
	post_visit_rows_count := types.AmountCreate(1000)
	post_edition_amount := types.AmountCreate(1000)

	shared.FixtureTimeMeasure(func() {
		user_bulker := shared.Bulker[model.User]{
			Amount_to_create: types.AmountCreate(max_users),
			Bulk_max:         types.BulkMax(8000),
			Dbname:           dbname,
		}
		user_bulker.Init().BulkCreate(func(u *model.User) { u.Fill() })

		post_bulker := shared.Bulker[model.Post]{
			Amount_to_create: types.AmountCreate(post_amount),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           dbname,
		}
		user_counter := NewCounter(int(max_users))
		post_bulker.Init().BulkCreate(func(p *model.Post) {
			p.Fill(user_counter.Next())
		})

		post_visit_bulker := shared.Bulker[model.PostVisits]{
			Amount_to_create: types.AmountCreate(post_visit_rows_count),
			Bulk_max:         types.BulkMax(16000),
			Dbname:           dbname,
		}
		post_counter := NewCounter(post_amount)
		post_visit_bulker.Init().BulkCreate(func(p *model.PostVisits) {
			p.Fill(post_counter.Next())
		})

		// TODO try to achieve high performance despite trigger for post_edition :/
		post_edition_bulker := shared.Bulker[model.PostEdition]{
			Amount_to_create: post_edition_amount,
			Bulk_max:         types.BulkMax(4000),
			Dbname:           dbname,
		}
		user_counter = NewCounter(int(max_users))
		post_counter = NewCounter(post_amount)
		post_edition_bulker.Init().BulkCreate(func(p *model.PostEdition) {
			p.Fill(post_counter.Next(), user_counter.Next())
		})

	}, "database filling with data")
}

func RunSubTests(task_number string, t *testing.T, test_func func(db_params testdb.DBParams)) {
	task_name := fmt.Sprintf("test%s", task_number)

	t.Run(task_name+"_unit_test", func(t *testing.T) {
		test_func(testdb.UnitTests)
	})

	if settings.ENABLED_PERFORMANCE_TESTS {
		t.Run(task_name+"_perf_with_index", func(t *testing.T) {
			shared.FixtureTimeMeasure(func() {
				test_func(testdb.PerformanceWithIndexes)
			}, task_name+"_perf_with_index")
		})
		t.Run(task_name+"_perf__without_index", func(t *testing.T) {
			shared.FixtureTimeMeasure(func() {
				test_func(testdb.PerformanceIndexless)
			}, task_name+"_perf__without_index")
		})
	}
}
