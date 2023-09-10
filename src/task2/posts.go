package task2

import (
	"math/rand"
	"time"
)

var (
	// Sequence counter which user to create
	PostIDSeq = 0
)

type Post struct {
	ID        int       `gorm:"column:id;primaryKey"`
	AuthorID  int       `gorm:"column:author_id"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

var Statuses []string = []string{"published", "draft", "archived"}

func (p Post) TableName() string {
	return "post"
}

func (p *Post) Fill(authorID int) {
	p.ID = GetNext(&PostIDSeq)
	p.AuthorID = authorID
	p.Status = Statuses[rand.Intn(3)]
	p.CreatedAt = time.Now()
}
