package processors

import (
	"fmt"
	"streamEtl/manager"
	"streamEtl/types"
	"time"
)

const transactionSize = 4

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
		//fmt.Printf("ResultLoader: Processing job ID %d\n", job.ID)

		// Process results in transactions
		transaction := make([]types.Result, 0, transactionSize)
		for _, result := range job.Result {
			transaction = append(transaction, result)

			if len(transaction) == transactionSize {
				processTransaction(transaction)
				transaction = transaction[:0]
			}
		}

		// Process remaining results if any
		if len(transaction) > 0 {
			processTransaction(transaction)
		}
	}
}

var transactionCounter int

func processTransaction(transaction []types.Result) {
	transactionCounter++
	fmt.Printf("\nResultLoader: Saving transaction #%d (%d results)\n", transactionCounter, len(transaction))
	fmt.Println("Results in this transaction:")
	for i, result := range transaction {
		fmt.Printf("  [%d] Result ID: %d, Job ID: %d\n",
			i+1, result.ResultID, result.JobID)
	}
	time.Sleep(30 * time.Millisecond)
}
