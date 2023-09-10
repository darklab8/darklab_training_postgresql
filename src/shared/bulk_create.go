package shared

import (
	"darklab_training_postgres/utils/types"
	"database/sql"
	"fmt"
	"math"

	"log/slog"

	"gorm.io/gorm"
)

type BulkJob[T any] struct {
	id     int
	done   bool
	Ptrs   []*T
	dbname types.Dbname
	result *gorm.DB
}

func (data *BulkJob[T]) runJob(worker_id int) StatusCode {
	FixtureTimeMeasure(func() {
		// fmt.Println("worker", worker_id, "started  job", data.id)
		FixtureConn(data.dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB) {
			data.result = conn_orm.Create(data.Ptrs)
		})
	}, fmt.Sprintf("worker %d finished job %d", worker_id, data.id))

	data.done = true
	return CodeSuccess
}

func BulkCreate[T any](
	dbname types.Dbname,
	amount_to_create types.AmountCreate,

	// Postgresql maximum parameters in one request 65535.
	// Proportional to amount of attributes
	bulk_max types.BulkMax,
	conn_orm *gorm.DB,
	fill func(*T),
) {
	job_timeout := 30
	worker_count := 10

	jobPool := JobPool[T, *BulkJob[T]]{
		JobTimeout: job_timeout,
		numWorkers: worker_count,
	}
	jobs := []*BulkJob[T]{}
	FixtureTimeMeasure(func() {

		bulk_times := int(math.Ceil(float64(amount_to_create) / float64(bulk_max)))
		left_to_create := int(amount_to_create)

		for i := 0; i < bulk_times; i++ {
			users := make([]T, bulk_max)
			usersPtrs := make([]*T, bulk_max)

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

			jobs = append(jobs, &BulkJob[T]{
				Ptrs:   usersPtrs,
				dbname: dbname,
				id:     i,
			})

			left_to_create -= creating_count
		}

		FixtureTimeMeasure(func() {
			jobPool.doJobs(jobs)
		}, "doing jobs")

		for job_number, job := range jobs {
			if !job.done {
				panic(fmt.Sprintf("job %d failed", job_number))
			}
			if job.result.Error != nil {
				slog.Error("failed to create users")
				panic(job.result.Error)
			}
			//fmt.Println("job ", job_number, "suceded. job_result=", job.result)
		}
	}, "BulkCreate whole")
}
