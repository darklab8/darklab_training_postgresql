package model

import (
	"darklab_training_postgres/golang/shared/utils"
	"math/rand"
	"time"

	"gorm.io/datatypes"
)

var (
	// Sequence counter which user to create
	PostIDSeq = 0
)

type Post struct {
	ID        int            `gorm:"column:id;primaryKey"`
	AuthorID  int            `gorm:"column:author_id"`
	Status    string         `gorm:"column:status"`
	CreatedAt datatypes.Date `gorm:"column:created_at"`
}

var Statuses []string = []string{"published", "draft", "archived"}

func (p Post) TableName() string {
	return "post"
}

func GerRandomTime() datatypes.Date {
	curTime := time.Now()
	curTime = curTime.Add(time.Hour * time.Duration(-rand.Intn(24*365)))
	return datatypes.Date(curTime)
}

func (p *Post) Fill(authorID int) {
	p.ID = utils.GetNext(&PostIDSeq)
	p.AuthorID = authorID
	p.Status = Statuses[rand.Intn(3)]
	p.CreatedAt = datatypes.Date(GerRandomTime())
}
