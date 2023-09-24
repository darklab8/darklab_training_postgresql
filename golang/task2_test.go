package golang

import (
	"darklab_training_postgres/golang/shared"
	"darklab_training_postgres/golang/shared/model"
	"darklab_training_postgres/golang/shared/types"
	"darklab_training_postgres/golang/testdb"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

func TestCreateData(t *testing.T) {
	shared.FixtureConn(testdb.UnitTests.Dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
		var count int64

		conn_orm.Model(&model.User{}).Count(&count)
		assert.Equal(t, int(testdb.UnitTests.MaxUsers), int(count))

		conn_orm.Model(&model.Post{}).Count(&count)
		assert.Equal(t, int(testdb.UnitTests.MaxUsers)*int(testdb.UnitTests.PostsPerUser), int(count))
	})
}
