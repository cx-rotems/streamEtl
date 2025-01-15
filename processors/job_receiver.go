package processors

import (
	"fmt"
	"streamEtl/manager"
	"streamEtl/types"
	"time"
)

// JobReceiver simulates getting jobs and sending them to MinioExtractor
type JobReceiver struct {
	jobChan    chan types.Job
	minioChan  chan types.Job
	jobManager *manager.JobManager
}

func NewJobReceiver(jobChan, minioChan chan types.Job, jm *manager.JobManager) *JobReceiver {
	return &JobReceiver{jobChan: jobChan, minioChan: minioChan, jobManager: jm}
}

func (jr *JobReceiver) Start() {
	defer jr.jobManager.WorkerDone()

	for job := range jr.jobChan {
		fmt.Printf("JobReceiver: Processing job ID %d\n", job.ID)
		time.Sleep(50 * time.Millisecond)
		jr.minioChan <- job
	}
	close(jr.minioChan)
}
