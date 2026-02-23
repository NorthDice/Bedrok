package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// ExecutionStatus represents the state of a single job run.
type ExecutionStatus string

const (
	// ExecutionStatusPending means the execution is queued but not yet picked up.
	ExecutionStatusPending ExecutionStatus = "pending"
	// ExecutionStatusRunning means a worker is currently processing this execution.
	ExecutionStatusRunning ExecutionStatus = "running"
	// ExecutionStatusCompleted means the execution finished successfully.
	ExecutionStatusCompleted ExecutionStatus = "completed"
	// ExecutionStatusFailed means the execution finished with an error.
	ExecutionStatusFailed ExecutionStatus = "failed"
)

// Execution records a single run of a Job.
type Execution struct {
	ID         uuid.UUID
	JobID      uuid.UUID
	Status     ExecutionStatus
	Attempt    int
	StartedAt  *time.Time
	FinishedAt *time.Time
	Result     json.RawMessage // output returned by the worker
	Error      string          // error message if status is failed
	CreatedAt  time.Time
}

// ExecutionResult holds the outcome reported by a worker after finishing a run.
type ExecutionResult struct {
	Result json.RawMessage
	Error  string
}
