package processors

import (
	"fmt"
	"streamEtl/types"
	"sync"
	"time"
)

const transactionSize = 4

type ResultLoader struct {
	loaderChan chan types.Job
	callback   func(int)
}

func NewResultLoader(loaderChan chan types.Job, callback func(int)) *ResultLoader {
	return &ResultLoader{loaderChan: loaderChan, callback: callback}
}

func (rl *ResultLoader) Start() {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
	}()

	for job := range rl.loaderChan {
		wg.Add(1)
		go func(job types.Job) {
			defer wg.Done()
			//fmt.Printf("ResultLoader: Processing job ID %d\n", job.ID)

			transaction := make([]types.Result, 0, transactionSize)
			for _, result := range job.Results {
				transaction = append(transaction, result)

				if len(transaction) == transactionSize {
					processTransaction(transaction)
					transaction = transaction[:0]
				}
			}

			if len(transaction) > 0 {
				processTransaction(transaction)
			}

			// Notify job completion
			rl.callback(job.ID)
		}(job)
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
