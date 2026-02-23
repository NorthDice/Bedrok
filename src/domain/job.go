package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// JobStatus represents the lifecycle state of a job.
type JobStatus string

const (
	// JobStatusActive means the job is enabled and will be scheduled.
	JobStatusActive JobStatus = "active"
	// JobStatusPaused means the job exists but will not be scheduled.
	JobStatusPaused JobStatus = "paused"
	// JobStatusDeleted marks the job as soft-deleted.
	JobStatusDeleted JobStatus = "deleted"
)

// Job represents a schedulable unit of work.
type Job struct {
	ID          uuid.UUID
	Name        string
	Payload     json.RawMessage // arbitrary data passed to the worker
	Schedule    string          // cron expression, e.g. "*/5 * * * *"; empty means one-time
	Status      JobStatus
	MaxAttempts int
	NextRunAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
