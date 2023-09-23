package shared

import (
	"context"
	"darklab_training_postgres/golang/shared/types"
	"database/sql"
	"fmt"
	"math"

	"github.com/uptrace/bun"
	"gorm.io/gorm"
)

type BulkJob[T any] struct {
	id     int
	done   bool
	Ptrs   []*T
	dbname types.Dbname
}

func (data *BulkJob[T]) runJob(worker_id int) StatusCode {
	FixtureTimeMeasure(func() {
		// fmt.Println("worker", worker_id, "started  job", data.id)
		FixtureConn(data.dbname, func(dbname types.Dbname, conn *sql.DB, conn_orm *gorm.DB, bundb *bun.DB) {
			_, err := bundb.NewInsert().Model(&data.Ptrs).Exec(context.TODO())
			if err != nil {
				panic(err)
			}
		})
	}, fmt.Sprintf("worker %d finished job %d", worker_id, data.id))

	data.done = true
	return CodeSuccess
}

type Bulker[T any] struct {
	all_objs    [][]T
	all_objPtrs [][]*T

	Amount_to_create types.AmountCreate
	// Postgresql maximum parameters in one request 65535.
	// Proportional to amount of attributes
	Bulk_max types.BulkMax

	bulk_times int
	Dbname     types.Dbname
}

func (b *Bulker[T]) Init() *Bulker[T] {
	b.bulk_times = int(math.Ceil(float64(b.Amount_to_create) / float64(b.Bulk_max)))

	b.all_objs = [][]T{}
	b.all_objPtrs = [][]*T{}
	for i := 0; i < b.bulk_times; i++ {
		b.all_objs = append(b.all_objs, make([]T, b.Bulk_max))
		b.all_objPtrs = append(b.all_objPtrs, make([]*T, b.Bulk_max))
	}
	return b
}

func (b *Bulker[T]) BulkCreate(
	fill func(*T),
) {
	job_timeout := 300
	worker_count := 10

	left_to_create := int(b.Amount_to_create)

	jobPool := JobPool[T, *BulkJob[T]]{
		JobTimeout: job_timeout,
		numWorkers: worker_count,
	}
	jobs := []*BulkJob[T]{}
	FixtureTimeMeasure(func() {

		FixtureTimeMeasure(func() {

			var objs []T
			var obj_ptrs []*T

			for i := 0; i < b.bulk_times; i++ {
				objs = b.all_objs[i]
				obj_ptrs = b.all_objPtrs[i]

				creating_count := int(b.Bulk_max)
				if left_to_create < int(b.Bulk_max) {
					objs = make([]T, left_to_create)
					obj_ptrs = make([]*T, left_to_create)
					creating_count = int(left_to_create)
				}

				for number, _ := range obj_ptrs {
					fill(&objs[number])
					obj_ptrs[number] = &objs[number]
				}

				jobs = append(jobs, &BulkJob[T]{
					Ptrs:   obj_ptrs,
					dbname: b.Dbname,
					id:     i,
				})

				left_to_create -= creating_count
			}
		}, "allocating memory")

		FixtureTimeMeasure(func() {
			jobPool.doJobs(jobs)
		}, "doing jobs")

		for job_number, job := range jobs {
			if !job.done {
				panic(fmt.Sprintf("job %d failed", job_number))
			}
			//fmt.Println("job ", job_number, "suceded. job_result=", job.result)
		}
	}, "BulkCreate whole")
}
