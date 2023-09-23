package model

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"table:post"`
	ID            *sql.NullInt32 `bun:"id,pk,autoincrement,scanonly"`
	AuthorID      int            `bun:"author_id"`
	Status        string         `bun:"status"`
	CreatedAt     time.Time      `bun:"created_at"`
	Rating        int            `bun:"rating"`
}

var Statuses []string = []string{"published", "draft", "archived"}

func (p Post) TableName() string {
	return "post"
}

func GerRandomTime() time.Time {
	curTime := time.Now()
	curTime = curTime.Add(time.Hour * time.Duration(-rand.Intn(24*365)))
	return curTime
}

func (p *Post) Fill(authorID int) {
	p.AuthorID = authorID
	p.Status = Statuses[rand.Intn(3)]
	p.CreatedAt = GerRandomTime()
	p.Rating = rand.Intn(100)
}
