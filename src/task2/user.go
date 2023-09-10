package task2

import (
	"fmt"
	"time"
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

func (u *User) Fill() {
	u.ID = GetNext(&UserIDSeq)
	u.FirstName = fmt.Sprintf("first_name%d", GetNext(&UserFirstNameSeq))
	u.SecondName = fmt.Sprintf("second_name%d", GetNext(&UserSecondNameSeq))
	u.Birth_date = time.Now()
	u.Email = fmt.Sprintf("email%d", GetNext(&UserEmailSeq))
	u.Password = fmt.Sprintf("password%d", GetNext(&UserPasswordSeq))
	u.Address = fmt.Sprintf("address%d", GetNext(&UserAddressSeq))
	u.Rating = 0
}
