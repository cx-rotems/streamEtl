package processors

import (
	//"fmt"
	"streamEtl/manager"
	"streamEtl/types"
	"time"
)

type MinioExtractor struct {
	minioChan  chan types.Job
	resultChan chan types.Job
	jobManager *manager.JobManager
}

func NewMinioExtractor(minioChan, resultChan chan types.Job, jm *manager.JobManager) *MinioExtractor {
	return &MinioExtractor{minioChan: minioChan, resultChan: resultChan, jobManager: jm}
}

func (me *MinioExtractor) Start() {
	defer close(me.resultChan)

	for job := range me.minioChan {
		//fmt.Printf("MinioExtractor: Extracting data for job ID %d\n", job.ID)
		for i := 0; i < 50; i++ {
			time.Sleep(100 * time.Millisecond) // simulate download from Minio
			job.Result = append(job.Result, types.Result{ResultID: i, JobID: job.ID})
		}
		me.resultChan <- job
	}
}
