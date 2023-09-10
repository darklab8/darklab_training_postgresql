package shared

import (
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	jobPool := JobPool[any, *JobTest]{}

	jobs := []*JobTest{}
	for job_id := 1; job_id <= 10; job_id++ {
		jobs = append(jobs, &JobTest{id: job_id})
	}

	status_codes := jobPool.doJobs(jobs)

	fmt.Println("results=", status_codes)
	for job_number, job := range jobs {
		if !job.done {
			fmt.Println("job ", job_number, "failed")
			continue
		}
		fmt.Println("job ", job_number, "suceded. job_result=", job.result)
	}
}
