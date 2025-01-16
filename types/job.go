package types

// Job represents a unit of work in the ETL pipeline
type Job struct {
	ID      int
	Results []Result
}

type Result struct {
	ResultID   int
	CvssScores string
	JobID      int
}
