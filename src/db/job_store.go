package db

import (
	"context"
	"time"

	"bedrok/domain"

	"github.com/google/uuid"
)

// JobFilter holds optional parameters for listing jobs.
type JobFilter struct {
	Status *domain.JobStatus
	Limit  int
	Offset int
}

// JobStore defines persistence operations for Job entities.
type JobStore interface {
	// Create persists a new job and sets its ID and timestamps.
	Create(ctx context.Context, job *domain.Job) error

	// GetByID returns a job by its unique identifier.
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Job, error)

	// List returns jobs matching the given filter.
	List(ctx context.Context, filter JobFilter) ([]*domain.Job, error)

	// Update saves changes to an existing job.
	Update(ctx context.Context, job *domain.Job) error

	// Delete soft-deletes a job by setting its status to deleted.
	Delete(ctx context.Context, id uuid.UUID) error

	// GetDue returns active jobs whose next_run_at is at or before now.
	// limit controls the maximum number of jobs returned in one call.
	GetDue(ctx context.Context, now time.Time, limit int) ([]*domain.Job, error)
}
