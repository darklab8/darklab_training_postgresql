package testdb

import (
	"darklab_training_postgres/golang/shared/types"
	"fmt"
	"strings"
)

type DBParams struct {
	Dbname        types.Dbname
	MaxUsers      types.MaxUsers
	PostsPerUser  types.PostsPerUser
	PostVisits    types.PostVisitsAmount
	PostEditions  types.PostEditionAmount
	PostApprovals types.PostApprovalAmount
}

func (d DBParams) ToStr() string {
	return strings.Join([]string{
		fmt.Sprintf("dbname=%s", string(d.Dbname)),
		fmt.Sprintf("users=%d", int(d.MaxUsers)),
		fmt.Sprintf("posts=%d", int(d.MaxUsers)*int(d.PostsPerUser)),
		fmt.Sprintf("post_visits=%d", int(d.PostVisits)),
		fmt.Sprintf("post_editions=%d", int(d.PostEditions)),
	}, " ")
}

var (
	UnitTests = DBParams{
		MaxUsers:      10000,
		PostsPerUser:  50,
		PostVisits:    1000,
		PostEditions:  1000,
		PostApprovals: 1000,
	}
	PerformanceIndexless = DBParams{
		MaxUsers:      50000,
		PostsPerUser:  50,
		PostVisits:    50000 * 10,
		PostEditions:  50000 * 10,
		PostApprovals: 10000,
	}
	PerformanceWithIndexes = DBParams{
		MaxUsers:     PerformanceIndexless.MaxUsers,
		PostsPerUser: PerformanceIndexless.PostsPerUser,
		PostVisits:   PerformanceIndexless.PostVisits,
		PostEditions: PerformanceIndexless.PostEditions,
	}
)
