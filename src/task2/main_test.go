package task2

import (
	"darklab_training_postgres/src/shared"
	"darklab_training_postgres/utils"
	"darklab_training_postgres/utils/types"
	"database/sql"
	_ "embed"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	//go:embed migrations/task2_1.sql
	Migration1 string

	//go:embed migrations/task2_2_disable_triggers_for_tests.sql
	Migration2 string
)

func FixtureTask2Migrations(conn *sql.DB) {
	fmt.Println()

	utils.MustExec(conn, Migration1)
	utils.MustExec(conn, Migration2)
}

// CREATE TABLE user_
// (
//   	id SERIAL PRIMARY KEY,
//   	first_name VARCHAR(100) NOT NULL,
// 	second_name VARCHAR(100) NOT NULL,
//   	birth_date DATE NOT NULL,
// 	email VARCHAR(100) NOT NULL,
// 	password VARCHAR(100) NOT NULL,
// 	address VARCHAR(200) NOT NULL
// );

type User struct {
	ID         int       `gorm:"column:id;primaryKey"`
	FirstName  string    `gorm:"column:first_name"`
	SecondName string    `gorm:"column:second_name"`
	Birth_date time.Time `gorm:"column:birth_date"`
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	Address    string    `gorm:"column:address"`
	Rating     int       `gorm:"column:rating"`
}

func (u User) TableName() string {
	return "user_"
}

const (
	PostgresqlSerialMax = 2147483647
)

var (
	// Sequence counter which user to create
	UserIDSeq         = 0
	UserFirstNameSeq  = 0
	UserSecondNameSeq = 0
	UserEmailSeq      = 0
	UserPasswordSeq   = 0
	UserAddressSeq    = 0
)

func GetNext(number *int) int {
	*number++
	return *number
}

func (u *User) fill() {
	u.ID = GetNext(&UserIDSeq)
	u.FirstName = fmt.Sprintf("first_name%d", GetNext(&UserFirstNameSeq))
	u.SecondName = fmt.Sprintf("second_name%d", GetNext(&UserSecondNameSeq))
	u.Birth_date = time.Now()
	u.Email = fmt.Sprintf("email%d", GetNext(&UserEmailSeq))
	u.Password = fmt.Sprintf("password%d", GetNext(&UserPasswordSeq))
	u.Address = fmt.Sprintf("address%d", GetNext(&UserAddressSeq))
	u.Rating = 0
}

const max_users = 10

func TestMain(t *testing.T) {
	shared.FixtureConnTestDB(func(dbname types.Dbname, conn *sql.DB) {
		FixtureTask2Migrations(conn)

		dsn := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=%s port=5432 sslmode=disable", dbname)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		users := make([]User, max_users)
		//#usersPtrs := [max_users]*User{}
		usersPtrs := make([]*User, max_users)

		for number, _ := range users {
			users[number].fill()
			usersPtrs[number] = &users[number]
		}
		result := db.Create(usersPtrs)
		assert.Nil(t, result.Error, "failed to create users")

		var user_count int
		rows, _ := conn.Query("SELECT count(*) FROM user_")
		rows.Next()
		rows.Scan(&user_count)
		assert.Equal(t, max_users, user_count)

	})
}
