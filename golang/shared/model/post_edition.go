package model

import (
	"darklab_training_postgres/golang/shared/model/array"
	"darklab_training_postgres/golang/shared/utils"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
)

var (
	// Sequence counter which user to create
	PostEditionTitleSeq   = 0
	PostEditionContentSeq = 0
)

type PostEdition struct {
	bun.BaseModel `bun:"table:post_edition"`

	ID       *sql.NullInt32 `bun:"id,pk,autoincrement,scanonly"`
	PostID   int            `bun:"post_id"`
	UserID   int            `bun:"user_id"`
	EditedAt time.Time      `bun:"edited_at"`
	Title    string         `bun:"title"`
	Content  string         `bun:"content"`
	Tags     []string       `bun:"tags,array"`
}

var (
	RandomTags = []string{"a1", "b1", "c1"}
)

func (p *PostEdition) Fill(postID int, UserID int) {
	p.PostID = postID
	p.UserID = UserID

	p.EditedAt = time.Time(GerRandomTime())
	p.Title = fmt.Sprintf("title%d", utils.GetNext(&PostEditionTitleSeq))
	p.Content = fmt.Sprintf("content%d", utils.GetNext(&PostEditionContentSeq))
	p.Tags = array.Array(append(p.Tags, RandomTags[rand.Intn(3)]))
}
