package golang_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"darklab_training_postgres/golang"
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestCount(t *testing.T) {

	counter := golang.NewCounter(3)

	assert.Equal(t, 1, counter.Next())
	assert.Equal(t, 2, counter.Next())
	assert.Equal(t, 3, counter.Next())
	assert.Equal(t, 1, counter.Next())
}

func TestInsertTags(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		golang.FixtureTask2Migrations(conn)

		user := model.User{}
		user.Fill()
		conn_orm.Create(&user)

		post := model.Post{}
		post.Fill(1)
		conn_orm.Create(&post)

		post_edition := model.PostEdition{}
		post_edition.Fill(1, 1)
		res, err := bundb.NewInsert().Model(&post_edition).Exec(context.TODO())
		_ = res
		utils.Check(err)

		post_editions := make([]model.PostEdition, 2)
		post_editions_ptrs := make([]*model.PostEdition, 2)

		for i, _ := range post_editions {
			post_editions[i].Fill(1, 1)
			post_editions_ptrs[i] = &post_editions[i]
		}
		res, err = bundb.NewInsert().Model(&post_editions_ptrs).Exec(context.TODO())
		_ = res
		utils.Check(err)
		golang.FixtureTask3Migrations(conn)

		fmt.Println("teardown")
	})
}

func TestInsertTags2(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		golang.FixtureTask2Migrations(conn)
		golang.FixtureFillWithData(
			dbname,
			types.MaxUsers(1000),
			types.PostsPerUser(5),
		)
		golang.FixtureTask3Migrations(conn)

		fmt.Println("teardown")
	})

}
