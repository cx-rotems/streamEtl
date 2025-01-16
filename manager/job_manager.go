package manager

import (
	"fmt"
)

// JobManager coordinates and tracks the completion of ETL processes
type JobManager struct {
	completionChan chan int
}

func NewJobManager() *JobManager {
	return &JobManager{
		completionChan: make(chan int, 100), // Buffer size to prevent blocking
	}
}

// Add new method to track job completion
func (jm *JobManager) JobCompleted(jobID int) {
	fmt.Printf("Job %d completed\n", jobID)
}

// Add new method to track new jobs
func (jm *JobManager) AddJob(jobID int) {
	fmt.Printf("Starting job %d\n", jobID)
}
