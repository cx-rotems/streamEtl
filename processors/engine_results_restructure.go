package processors

import (
	"fmt"
	"streamEtl/types"
	"sync"
	"time"
)

type EngineResultsRestructure struct {
	resultChan     chan types.Job
	enrichmentChan chan types.Job
}

func NewEngineResultsRestructure(resultChan, enrichmentChan chan types.Job) *EngineResultsRestructure {
	return &EngineResultsRestructure{resultChan: resultChan, enrichmentChan: enrichmentChan}
}

func (er *EngineResultsRestructure) Start() {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		close(er.enrichmentChan)
	}()

	for job := range er.resultChan {
		wg.Add(1)
		go func(job types.Job) {
			defer wg.Done()
			for i := 0; i < len(job.Results); i++ {
				job.Results[i].CvssScores = fmt.Sprintf("%d", i*10)
				time.Sleep(70 * time.Millisecond) // simulate restructure
				//fmt.Printf("EngineResultsRestructure: Restructuring result for result ID %d and job ID  %d\n", job.Results[i].ResultID, job.Results[i].JobID)
			}
			er.enrichmentChan <- job
		}(job)
	}
}
