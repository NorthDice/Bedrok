-- +goose Up
CREATE TYPE IF NOT EXISTS bedrok.execution_status AS ENUM ('pending', 'running', 'completed', 'failed');

CREATE TABLE IF NOT EXISTS bedrok.executions (
    id          UUID                    PRIMARY KEY DEFAULT gen_random_uuid(),
    job_id      UUID                    NOT NULL REFERENCES bedrok.jobs (id) ON DELETE CASCADE,
    status      bedrok.execution_status NOT NULL DEFAULT 'pending',
    attempt     INT                     NOT NULL DEFAULT 1,
    started_at  TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    result      JSONB,
    error       TEXT,
    created_at  TIMESTAMPTZ             NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_executions_job_id ON bedrok.executions (job_id);
CREATE INDEX IF NOT EXISTS idx_executions_status ON bedrok.executions (status);

-- +goose Down
DROP TABLE IF EXISTS bedrok.executions;
DROP TYPE  IF EXISTS bedrok.execution_status;
