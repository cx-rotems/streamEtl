package processors

import (
	"fmt"
	"streamEtl/manager"
	"streamEtl/types"
	"time"
)

type ResultLoader struct {
	loaderChan chan types.Job
	jobManager *manager.JobManager
}

func NewResultLoader(loaderChan chan types.Job, jm *manager.JobManager) *ResultLoader {
	return &ResultLoader{loaderChan: loaderChan, jobManager: jm}
}

func (rl *ResultLoader) Start() {
	defer rl.jobManager.WorkerDone()

	for job := range rl.loaderChan {
		fmt.Printf("ResultLoader: Loading result for job ID %d\n", job.ID)
		for i := 0; i < len(job.Result); i++ {
			fmt.Printf("ResultLoader: Saving result ID %d for job ID %d\n", job.Result[i].ResultID, job.ID)
		}
		time.Sleep(30 * time.Millisecond)
	}
}
