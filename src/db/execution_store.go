package db

import (
	"context"

	"bedrok/domain"

	"github.com/google/uuid"
)

// ExecutionStore defines persistence operations for Execution entities.
type ExecutionStore interface {
	// Create persists a new execution record.
	Create(ctx context.Context, exec *domain.Execution) error

	// GetByID returns an execution by its unique identifier.
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Execution, error)

	// ListByJob returns the most recent executions for the given job.
	ListByJob(ctx context.Context, jobID uuid.UUID, limit int) ([]*domain.Execution, error)

	// UpdateStatus updates the status and result of an execution.
	UpdateStatus(ctx context.Context, id uuid.UUID, status domain.ExecutionStatus, result *domain.ExecutionResult) error
}
