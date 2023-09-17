package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	Query1 string
	Query2 string
	Query3 string
	Query4 string
	Query5 string
)

func init() {
	Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_1.sql"))
	Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_2.sql"))
	Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_3.sql"))
	Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_4.sql"))
	Query5 = utils.GetSQLFile(utils.ReadProjectFile("sql/task3/queries/query3_5.sql"))
}

func TestQueryReuseSetup2(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		var count int
		rows, _ := conn.Query("SELECT count(*) FROM post")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, int(TempDb.MaxUsers)*int(TempDb.PostsPerUser), count)
	})
}

func TestQuery1(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		author_id := 1 + rand.Intn(int(TempDb.MaxUsers)-1)
		result := conn_orm.Raw(Query1, sql.Named("author_id", author_id))
		// rows, err := conn.Query(Query1, row)
		if result.Error != nil {
			panic(result.Error)
		}

		var count int
		result.Scan(&count)
		assert.Equal(t, int(TempDb.PostsPerUser), count)
	})
}

func TestQuery2(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		N := rand.Intn(int(TempDb.MaxUsers) + int(TempDb.PostsPerUser))
		result := conn_orm.Raw(Query2, sql.Named("N", N))
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

func TestQuery3(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		N := (rand.Intn(int(TempDb.MaxUsers)+int(TempDb.PostsPerUser)) / 4)
		result := conn_orm.Raw(Query3, sql.Named("N", N))
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

func TestQuery4(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		N := (rand.Intn(int(TempDb.MaxUsers)+int(TempDb.PostsPerUser)) / 4)
		result := conn_orm.Raw(Query3, sql.Named("N", N))
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

func TestMigration(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		FixtureTask2Migrations(conn)
		FixtureTask3Migrations(conn_orm)
	})
}
