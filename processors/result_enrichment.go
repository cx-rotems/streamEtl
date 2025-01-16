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
	defer close(re.loaderChan)

	for job := range re.enrichmentChan {
		for i := 0; i < len(job.Results); i++ {
			job.Results[i].CvssScores = job.Results[i].CvssScores + " enrichment"
			time.Sleep(60 * time.Millisecond) // simulate result enrichment
			//fmt.Printf("ResultEnrichment: Enriching result for result ID %d and job ID  %d\n",  job.Results[i].ResultID,  job.Results[i].JobID) // simulate result enrichment
		}
		re.loaderChan <- job
	}
}
