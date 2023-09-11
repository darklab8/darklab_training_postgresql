package orm

import (
	"darklab_training_postgres/golang/shared/utils"
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
	u.ID = utils.GetNext(&UserIDSeq)
	u.FirstName = fmt.Sprintf("first_name%d", utils.GetNext(&UserFirstNameSeq))
	u.SecondName = fmt.Sprintf("second_name%d", utils.GetNext(&UserSecondNameSeq))
	u.Birth_date = time.Now()
	u.Email = fmt.Sprintf("email%d", utils.GetNext(&UserEmailSeq))
	u.Password = fmt.Sprintf("password%d", utils.GetNext(&UserPasswordSeq))
	u.Address = fmt.Sprintf("address%d", utils.GetNext(&UserAddressSeq))
	u.Rating = 0
}
