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
		fmt.Printf("EngineResultsRestructure: Restructuring result for job ID %d\n", job.ID)
		time.Sleep(70 * time.Millisecond)
		er.enrichmentChan <- job
	}
	close(er.enrichmentChan)
}
