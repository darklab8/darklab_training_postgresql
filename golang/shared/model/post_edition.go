package model

import (
	"darklab_training_postgres/golang/shared/model/array"
	"darklab_training_postgres/golang/shared/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/uptrace/bun"
	"gorm.io/datatypes"
)

var (
	// Sequence counter which user to create
	PostEditionIDSeq      = 0
	PostEditionTitleSeq   = 0
	PostEditionContentSeq = 0
)

type PostEdition struct {
	bun.BaseModel `bun:"table:post_edition"`

	ID       int       `bun:"id,pk"`
	PostID   int       `bun:"post_id"`
	UserID   int       `bun:"user_id"`
	EditedAt time.Time `bun:"edited_at"`
	Title    string    `bun:"title"`
	Content  string    `bun:"content"`
	Tags     []string  `bun:"tags,array"`
}

var (
	RandomTags = []string{"a1", "b1", "c1"}
)

func (p *PostEdition) Fill(postID int, UserID int) {
	p.ID = utils.GetNext(&PostEditionIDSeq)
	p.PostID = postID
	p.UserID = UserID

	p.EditedAt = time.Time(GerRandomTime())
	p.Title = fmt.Sprintf("title%d", utils.GetNext(&PostEditionTitleSeq))
	p.Content = fmt.Sprintf("content%d", utils.GetNext(&PostEditionContentSeq))
	p.Tags = array.Array(append(p.Tags, RandomTags[rand.Intn(3)]))
}

type PostEditionGorm struct {
	ID       int            `gorm:"column:id;primaryKey"`
	PostID   int            `gorm:"column:post_id"`
	UserID   int            `gorm:"column:user_id"`
	EditedAt datatypes.Date `gorm:"column:edited_at"`
	Title    string         `gorm:"column:title"`
	Content  string         `gorm:"column:content"`
}

func (p PostEditionGorm) TableName() string {
	return "post_edition"
}

func (p *PostEditionGorm) Fill(postID int, UserID int) {
	p.ID = utils.GetNext(&PostEditionIDSeq)
	p.PostID = postID
	p.UserID = UserID

	p.EditedAt = (GerRandomTime())
	p.Title = fmt.Sprintf("title%d", utils.GetNext(&PostEditionTitleSeq))
	p.Content = fmt.Sprintf("content%d", utils.GetNext(&PostEditionContentSeq))
}
