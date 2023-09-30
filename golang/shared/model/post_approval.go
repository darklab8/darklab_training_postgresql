package model

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
)

type PostApproval struct {
	bun.BaseModel `bun:"table:post_approval"`

	ID        *sql.NullInt32 `bun:"id,pk,autoincrement,scanonly"`
	PostID    int            `bun:"post_id"`
	UserID    int            `bun:"user_id"`
	CreatedAt time.Time      `bun:"created_at"`
	Change    int            `bun:"change"`
}

var (
	ChangeChoices = []int{-1, 1}
)

func (p *PostApproval) Fill(postID int, UserID int) {
	p.PostID = postID
	p.UserID = UserID

	p.CreatedAt = time.Time(GerRandomTime())
	p.Change = ChangeChoices[rand.Intn(2)]
}
