package golang

import (
	"context"
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

var (
	Task4Query1   string
	Task4Query1_2 string
	Task4Query2   string
	Task4Query3   string
	Task4Query4   string
	Task4Query5   string
	Task4Query6   string
)

func init() {
	Task4Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_1.sql"))
	Task4Query1_2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_1_2.sql"))
	Task4Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_2.sql"))
	Task4Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_3.sql"))
	Task4Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_4.sql"))
	Task4Query5 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_5.sql"))
	Task4Query6 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_6.sql"))
}

func CountRows(result *gorm.DB) int {
	rows, err := result.Rows()
	if err != nil {
		panic(err)
	}
	count_rows := 0
	for rows.Next() {
		count_rows += 1
	}
	return count_rows
}

func TestTask4Query1MostVisitedPostInAYear(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

		Query1Test := func(query1 string) {
			N := 50
			result := conn_orm.Raw(query1, sql.Named("N", N))
			utils.Check(result.Error)

			count := CountRows(result)
			assert.Equal(t, int(TempDb.PostsPerUser), count)
		}

		t.Run("taskquery4_1", func(t *testing.T) {
			Query1Test(Task4Query1)
		})
		t.Run("taskquery4_2", func(t *testing.T) {
			Query1Test(Task4Query1_2)
		})
	})
}

type SaveNeverTestException struct{}

func (m SaveNeverTestException) Error() string {
	return "boom"
}

func TestTask4Query2MostVisitedPostsForAuthor(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		N := 50
		user_id := 10

		post_eds := make([]model.PostEdition, 100)
		post_ed_ptrs := make([]*model.PostEdition, 100)
		for i, _ := range post_ed_ptrs {
			post_eds[i].Fill(rand.Intn(10)+1, user_id)
			post_ed_ptrs[i] = &post_eds[i]
		}
		_, err := bundb.NewInsert().Model(&post_ed_ptrs).Exec(context.TODO())
		utils.Check(err)

		post_visits := make([]model.PostVisits, 100)
		post_visits_ptrs := make([]*model.PostVisits, 100)
		for i, _ := range post_visits_ptrs {
			post_visits[i].Fill(rand.Intn(10) + 1)
			post_visits_ptrs[i] = &post_visits[i]
		}
		_, err = bundb.NewInsert().On("CONFLICT DO NOTHING").Model(&post_visits_ptrs).Exec(context.TODO())
		utils.Check(err)

		Query1Test := func(query1 string) {

			result := conn_orm.Raw(
				query1,
				sql.Named("N", N),
				sql.Named("user_id", user_id),
			)
			utils.Check(result.Error)

			count := CountRows(result)
			assert.GreaterOrEqual(t, N, count)
		}

		t.Run("taskquery4_2", func(t *testing.T) {
			Query1Test(Task4Query2)
		})
	})
}

func TestTask4Query3(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		N := 50
		result := conn_orm.Raw(
			Task4Query3,
			sql.Named("N", N),
		)
		utils.Check(result.Error)

		count := CountRows(result)
		assert.Equal(t, N, count)
	})
}

func TestTask4Query4(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		N := 50
		result := conn_orm.Raw(
			Task4Query4,
			sql.Named("N", N),
			sql.Named("K", 2),
		)
		utils.Check(result.Error)

		count := CountRows(result)
		assert.Equal(t, N, count)
	})
}

func TestTask4Query5(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		N := 50
		result := conn_orm.Raw(
			Task4Query5,
			sql.Named("N", N),
		)
		utils.Check(result.Error)

		count := CountRows(result)
		assert.Equal(t, N, count)
	})
}

func TestTask4Query6(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		N := 50
		result := conn_orm.Raw(
			Task4Query6,
			sql.Named("N", N),
		)
		utils.Check(result.Error)

		tag_count := 3
		count := CountRows(result)
		assert.Equal(t, tag_count, count)
	})
}
