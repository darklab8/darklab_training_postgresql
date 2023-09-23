package model

import (
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

var (
	// Sequence counter which user to create
	UserFirstNameSeq  = 0
	UserSecondNameSeq = 0
	UserEmailSeq      = 0
	UserPasswordSeq   = 0
	UserAddressSeq    = 0
)

type User struct {
	bun.BaseModel `bun:"table:user_"`
	ID            *sql.NullInt32 `gorm:"id,pk,autoincrement"`
	FirstName     string         `gorm:"first_name"`
	SecondName    string         `gorm:"second_name"`
	Birth_date    time.Time      `gorm:"birth_date"`
	Email         string         `gorm:"email"`
	Password      string         `gorm:"password"`
	Address       string         `gorm:"address"`
	Rating        int            `gorm:"rating"`
}

func (u User) TableName() string {
	return "user_"
}

func (u *User) Fill() {
	u.FirstName = fmt.Sprintf("first_name%d", utils.GetNext(&UserFirstNameSeq))
	u.SecondName = fmt.Sprintf("second_name%d", utils.GetNext(&UserSecondNameSeq))
	u.Birth_date = GerRandomTime()
	u.Email = fmt.Sprintf("email%d", utils.GetNext(&UserEmailSeq))
	u.Password = fmt.Sprintf("password%d", utils.GetNext(&UserPasswordSeq))
	u.Address = fmt.Sprintf("address%d", utils.GetNext(&UserAddressSeq))
	u.Rating = 0
}
