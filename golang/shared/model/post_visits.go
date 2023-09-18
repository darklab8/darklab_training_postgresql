package model

import (
	"darklab_training_postgres/golang/shared/utils"
	"math/rand"

	"gorm.io/datatypes"
)

var (
	// Sequence counter which user to create
	PostVisitIDSeq = 0
)

type PostVisits struct {
	ID      int            `gorm:"column:id;primaryKey"`
	PostID  int            `gorm:"column:post_id"`
	DayDate datatypes.Date `gorm:"column:day_date"`
	Visits  int            `gorm:"column:visits"`
}

func (p PostVisits) TableName() string {
	return "post_visits_per_day"
}

func (p *PostVisits) Fill(postID int) {
	p.ID = utils.GetNext(&PostVisitIDSeq)
	p.PostID = postID
	p.DayDate = GerRandomTime()
	p.Visits = rand.Intn(1000)
}
