package golang

import (
	"context"
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestTask4Query2MostVisitedPostsForAuthor(t *testing.T) {
	RunSubTests("4_2", t, func(db_params testdb.DBParams) {
		shared.FixtureConn(db_params.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
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
	})
}
