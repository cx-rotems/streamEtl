package manager

import (
	"fmt"
	"sync"
)

// JobManager coordinates and tracks the completion of ETL processes
type JobManager struct {
	wg       sync.WaitGroup
	doneChan chan struct{}
	// Add job tracking
	jobsMutex  sync.Mutex
	activeJobs map[int]bool
}

func NewJobManager() *JobManager {
	return &JobManager{
		doneChan:   make(chan struct{}),
		activeJobs: make(map[int]bool),
	}
}

func (jm *JobManager) AddWorker() {
	jm.wg.Add(1)
}

func (jm *JobManager) WorkerDone() {
	jm.wg.Done()
}

func (jm *JobManager) WaitForCompletion() {
	jm.wg.Wait()
	close(jm.doneChan)
}

// Add new method to track job completion
func (jm *JobManager) JobCompleted(jobID int) {
	jm.jobsMutex.Lock()
	delete(jm.activeJobs, jobID)
	fmt.Printf("Job %d completed\n", jobID)
	jm.jobsMutex.Unlock()
}

// Add new method to track new jobs
func (jm *JobManager) AddJob(jobID int) {
	jm.jobsMutex.Lock()
	jm.activeJobs[jobID] = true
	jm.jobsMutex.Unlock()
}
