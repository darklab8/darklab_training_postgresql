package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateData(t *testing.T) {
	shared.FixtureConn(TempDb.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
		var count int64

		conn_orm.Model(&model.User{}).Count(&count)
		assert.Equal(t, int(TempDb.MaxUsers), int(count))

		conn_orm.Model(&model.Post{}).Count(&count)
		assert.Equal(t, int(TempDb.MaxUsers)*int(TempDb.PostsPerUser), int(count))
	})
}
