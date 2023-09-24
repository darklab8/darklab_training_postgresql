package testdb

import "darklab_training_postgres/golang/shared/types"

type DBParams struct {
	MaxUsers     types.MaxUsers
	PostsPerUser types.PostsPerUser
	Dbname       types.Dbname
}

var (
	UnitTests = DBParams{
		MaxUsers:     10000,
		PostsPerUser: 50,
	}
	PerformanceIndexless = DBParams{
		MaxUsers:     100000,
		PostsPerUser: 50,
	}
	PerformanceWithIndexes = DBParams{
		MaxUsers:     PerformanceIndexless.MaxUsers,
		PostsPerUser: PerformanceIndexless.PostsPerUser,
	}
)
