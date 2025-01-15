package processors

import (
	//"fmt"
	"streamEtl/manager"
	"streamEtl/types"
	"time"
)

type ResultEnrichment struct {
	enrichmentChan chan types.Job
	loaderChan     chan types.Job
	jobManager     *manager.JobManager
}

func NewResultEnrichment(enrichmentChan, loaderChan chan types.Job, jm *manager.JobManager) *ResultEnrichment {
	return &ResultEnrichment{enrichmentChan: enrichmentChan, loaderChan: loaderChan, jobManager: jm}
}

func (re *ResultEnrichment) Start() {
	defer re.jobManager.WorkerDone()

	for job := range re.enrichmentChan {
		for i := 0; i < len(job.Result); i++ {
			time.Sleep(60 * time.Millisecond) // simulate result enrichment
			//fmt.Printf("ResultEnrichment: Enriching result for result ID %d and job ID  %d\n",  job.Result[i].ResultID,  job.Result[i].JobID) // simulate result enrichment
		}
		re.loaderChan <- job
	}
	close(re.loaderChan)
}
