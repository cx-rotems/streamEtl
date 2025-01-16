package manager

import (
	"fmt"
	"sync"
)

// JobManager coordinates and tracks the completion of ETL processes
type JobManager struct {
	jobsMutex  sync.Mutex
	activeJobs map[int]bool
}

func NewJobManager() *JobManager {
	return &JobManager{
		activeJobs: make(map[int]bool),
	}
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
