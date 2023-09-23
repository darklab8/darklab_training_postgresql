package model

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
)

type PostVisits struct {
	bun.BaseModel `bun:"table:post_visits_per_day"`
	ID            *sql.NullInt32 `bun:"id,pk,autoincrement,scanonly"`
	PostID        int            `bun:"post_id"`
	DayDate       time.Time      `bun:"day_date"`
	Visits        int            `bun:"visits"`
}

func (p *PostVisits) Fill(postID int) {
	p.PostID = postID
	p.DayDate = GerRandomTime()
	p.Visits = rand.Intn(1000)
}
