package shared

import (
	"fmt"
	"time"
)

// ======================
// Test Example

type JobTest struct {
	// any desired arbitary data
	id     int
	result int
	done   bool
}

func (data *JobTest) runJob(worker_id int) StatusCode {
	fmt.Println("worker", worker_id, "started  job", data.id)
	time.Sleep(time.Second * time.Duration(data.id))
	fmt.Println("worker", worker_id, "finished job", data.id)
	data.result = data.id * 1
	data.done = true
	return CodeSuccess
}

// ====================

type IJob[T any] interface {
	// *JobTest | *JobDropRights | *JobGrantRights
	*JobTest | *BulkJob[T]
	runJob(worker_id int) StatusCode
}

type StatusCode int

const (
	CodeSuccess StatusCode = 0
	CodeTimeout StatusCode = 1
)

type JobPool[T any, jobd IJob[T]] struct {
	JobTimeout int // seconds
	numWorkers int
}

func (j JobPool[any, jobd]) launchWorker(worker_id int, jobs <-chan jobd, results chan<- StatusCode) {
	// fmt.Printf("worker %d started\n", worker_id)
	for job := range jobs {
		results <- job.runJob(worker_id)
	}
	// fmt.Printf("worker %d finished\n", worker_id)
}

func (j JobPool[any, jobd]) doJobs(jobs []jobd) []StatusCode {
	numJobs := len(jobs)

	// In order to use our pool of workers we need to send them work and collect their results.
	// We make 2 channels for this.
	jobs_channel := make(chan jobd, numJobs)
	result_channel := make(chan StatusCode, numJobs)

	status_codes := []StatusCode{}

	// This starts up N workers, initially blocked because there are no jobs yet.
	numWorker := 3
	if j.numWorkers != 0 {
		numWorker = j.numWorkers
	}
	for worker_id := 1; worker_id <= numWorker; worker_id++ {
		go j.launchWorker(worker_id, jobs_channel, result_channel)
	}

	// Here we send 5 jobs and
	for _, job := range jobs {
		jobs_channel <- job
	}
	// then close that channel to indicate that is all the work we have.
	close(jobs_channel)

	// added timeout
	jobTimeout := 3
	if j.JobTimeout != 0 {
		jobTimeout = j.JobTimeout
	}

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have finished.
	// An alternative way to wait for multiple goroutines is to use a WaitGroup.
	for job_number, _ := range jobs {
		select {
		case res := <-result_channel:
			status_codes = append(status_codes, res)
		case <-time.After(time.Duration(jobTimeout) * time.Second):
			// non zero exit by timeout
			status_codes = append(status_codes, CodeTimeout)
			fmt.Println("timeout for job_number=", job_number)
		}

	}
	return status_codes
}
