package core

import "fmt"

// Scheduler handles task scheduling
type Scheduler struct {}

// NewScheduler creates a new scheduler instance
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// ScheduleJob schedules a job to run at a specific time
func (s *Scheduler) ScheduleJob(job func(), delayInSeconds int) {
	// Add job scheduling logic here
	fmt.Printf("Job scheduled to run in %d seconds\n", delayInSeconds)
}
