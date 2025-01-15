package processors

import (
	"fmt"
	"streamEtl/manager"
	"streamEtl/types"
	"time"
)

type EngineResultsRestructure struct {
	resultChan     chan types.Job
	enrichmentChan chan types.Job
	jobManager     *manager.JobManager
}

func NewEngineResultsRestructure(resultChan, enrichmentChan chan types.Job, jm *manager.JobManager) *EngineResultsRestructure {
	return &EngineResultsRestructure{resultChan: resultChan, enrichmentChan: enrichmentChan, jobManager: jm}
}

func (er *EngineResultsRestructure) Start() {
	defer er.jobManager.WorkerDone()

	for job := range er.resultChan {
		for i := 0; i < len(job.Result); i++ {
			job.Result[i].CvssScores = fmt.Sprintf("%d", i*10)
			time.Sleep(70 * time.Millisecond) // simulate restructure
			fmt.Printf("EngineResultsRestructure: Restructuring result for result ID %d and job ID  %d\n", job.Result[i].ResultID, job.Result[i].JobID)
		}
		er.enrichmentChan <- job
	}
	close(er.enrichmentChan)
}
