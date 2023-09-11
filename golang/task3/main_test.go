package task3

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"darklab_training_postgres/golang/task2"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	//go:embed queries/query3_1.sql
	Query1 string

	//go:embed queries/query3_2.sql
	Query2 string

	//go:embed queries/query3_3.sql
	Query3 string

	//go:embed queries/query3_4.sql
	Query4 string

	//go:embed queries/query3_5.sql
	Query5 string
)

func FixtureDefaultData(dbname types.Dbname, conn *sql.DB) {
	task2.FixtureTask2Migrations(conn)

	task2.FixtureFillWithData(
		dbname,
		temporal_max_users,
		temporal_posts_per_users,
	)
}

var temporal_max_users types.MaxUsers = 1000
var temporal_posts_per_users types.PostsPerUser = 5
var temporal_dbname types.Dbname

func TestMain(m *testing.M) {
	fmt.Println("seting")
	var code int
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		FixtureDefaultData(dbname, conn)
		temporal_dbname = dbname
		code = m.Run()
		fmt.Println("teardown")
	})
	os.Exit(code)
}

func TestQueryReuseSetup1(t *testing.T) {
	shared.FixtureConn(temporal_dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		var count int
		rows, _ := conn.Query("SELECT count(*) FROM user_")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, count, 1000)
	})
}

func TestQueryReuseSetup2(t *testing.T) {
	shared.FixtureConn(temporal_dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		var count int
		rows, _ := conn.Query("SELECT count(*) FROM post")
		rows.Next()
		rows.Scan(&count)
		assert.Equal(t, count, 1000*5)
	})
}

func TestQuery1(t *testing.T) {
	shared.FixtureConn(temporal_dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		author_id := 1 + rand.Intn(int(temporal_max_users)-1)
		result := conn_orm.Raw(Query1, sql.Named("author_id", author_id))
		// rows, err := conn.Query(Query1, row)
		if result.Error != nil {
			panic(result.Error)
		}

		var count int
		result.Scan(&count)
		assert.Equal(t, int(temporal_posts_per_users), count)
	})
}

func TestQuery2(t *testing.T) {
	shared.FixtureConn(temporal_dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		N := rand.Intn(int(temporal_max_users) + int(temporal_posts_per_users))
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
	shared.FixtureConn(temporal_dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		N := (rand.Intn(int(temporal_max_users)+int(temporal_posts_per_users)) / 4)
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
	shared.FixtureConn(temporal_dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {

		fmt.Println()
		N := (rand.Intn(int(temporal_max_users)+int(temporal_posts_per_users)) / 4)
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
