package processors

import (
    "streamEtl/types"
    "sync"
    "time"
)

type MinioExtractor struct {
    minioChan  chan types.Job
    resultChan chan types.Job
}

func NewMinioExtractor(minioChan, resultChan chan types.Job) *MinioExtractor {
    return &MinioExtractor{minioChan: minioChan, resultChan: resultChan}
}

func (me *MinioExtractor) Start() {
    var wg sync.WaitGroup
    defer func() {
        wg.Wait()
        close(me.resultChan)
    }()

    for job := range me.minioChan {
        wg.Add(1)
        go func(job types.Job) {
            defer wg.Done()
            //fmt.Printf("MinioExtractor: Extracting data for job ID %d\n", job.ID)
            for i := 0; i <= 50; i++ {
                time.Sleep(100 * time.Millisecond) // simulate download from Minio
                job.Results = append(job.Results, types.Result{ResultID: i, JobID: job.ID})
            }
            me.resultChan <- job
        }(job)
    }
}
