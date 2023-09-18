package golang

import (
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
	Task3Query1 string
	Task3Query2 string
	Task3Query3 string
	Task3Query4 string
	Task3Query5 string
)

func init() {
	Task3Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_1.sql"))
	Task3Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_2.sql"))
	Task3Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_3.sql"))
	Task3Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_4.sql"))
	Task3Query5 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_5.sql"))
}

func TestQueryReuseSetup2(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		var count int64
		conn_orm.Model(&model.Post{}).Count(&count)
		assert.Equal(t, int(TempDb.MaxUsers)*int(TempDb.PostsPerUser), int(count))
	})
}

func TestQuery1UserPostCount(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		author_id := 1 + rand.Intn(int(TempDb.MaxUsers)-1)
		result := conn_orm.Raw(Task3Query1, sql.Named("author_id", author_id))
		if result.Error != nil {
			panic(result.Error)
		}

		var count int
		result.Scan(&count)
		assert.Equal(t, int(TempDb.PostsPerUser), count)
	})
}

func TestQuery2PublishedOrderedPosts(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		N := rand.Intn(int(TempDb.MaxUsers) + int(TempDb.PostsPerUser))
		result := conn_orm.Raw(Task3Query2, sql.Named("N", N))
		if result.Error != nil {
			panic(result.Error)
		}

		rows, err := result.Rows()
		if err != nil {
			panic(err)
		}

		count_rows := 0
		for rows.Next() {
			count_rows += 1
		}
		assert.Equal(t, N, count_rows)
	})
}

func TestQuery3PostsAwaitingPublishing(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

		N := (rand.Intn(int(TempDb.MaxUsers)+int(TempDb.PostsPerUser)) / 4)
		result := conn_orm.Raw(Task3Query3, sql.Named("N", N))
		if result.Error != nil {
			panic(result.Error)
		}

		rows, err := result.Rows()
		if err != nil {
			panic(err)
		}

		count_rows := 0
		for rows.Next() {
			count_rows += 1
		}
		assert.Equal(t, N, count_rows)
	})
}

func TestQuery4RecentlyUpdatedPostsByTag(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

		raw_sql := conn_orm.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Raw(Task3Query4,
				sql.Named("N", 20),
				sql.Named("K", 1),
				sql.Named("L", 10),
				sql.Named("tag", "a1"),
			)
		})
		result, err := conn.Exec(raw_sql)

		utils.Check(err)

		count_rows, err := result.RowsAffected()
		utils.Check(err)

		assert.Equal(t, 20, count_rows)
	})
}

func TestQuery5FindNPostsWithMostRating(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {

		result := conn_orm.Raw(Task3Query5,
			sql.Named("N", 20),
		)

		utils.Check(result.Error)
		count_rows := CountRows(result)

		assert.Equal(t, 20, count_rows)
	})
}

func TestMigration(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		FixtureTask2Migrations(conn)
		FixtureTask3Migrations(conn_orm, bundb)
	})
}
