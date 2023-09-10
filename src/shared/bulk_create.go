package shared

import (
	"darklab_training_postgres/utils/types"
	"math"

	"log/slog"

	"gorm.io/gorm"
)

func BulkCreate[T any](
	amount_to_create types.AmountCreate,

	// Postgresql maximum parameters in one request 65535.
	// Proportional to amount of attributes
	bulk_max types.BulkMax,
	conn_orm *gorm.DB,
	fill func(*T),
) {
	bulk_times := int(math.Ceil(float64(amount_to_create) / float64(bulk_max)))
	left_to_create := int(amount_to_create)

	users := make([]T, bulk_max)
	usersPtrs := make([]*T, bulk_max)

	for i := 0; i < bulk_times; i++ {
		creating_count := int(bulk_max)
		if left_to_create < int(bulk_max) {
			users = make([]T, left_to_create)
			usersPtrs = make([]*T, left_to_create)
			creating_count = int(left_to_create)
		}

		for number, _ := range usersPtrs {
			fill(&users[number])
			usersPtrs[number] = &users[number]
		}
		result := conn_orm.Create(usersPtrs)
		if result.Error != nil {
			slog.Error("failed to create users")
			panic(result.Error)
		}

		left_to_create -= creating_count
	}
}
