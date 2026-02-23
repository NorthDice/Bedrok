-- +goose Up
CREATE TYPE IF NOT EXISTS bedrok.job_status AS ENUM ('active', 'paused', 'deleted');

CREATE TABLE IF NOT EXISTS bedrok.jobs (
    id           UUID             PRIMARY KEY DEFAULT gen_random_uuid(),
    name         VARCHAR(255)     NOT NULL,
    payload      JSONB            NOT NULL DEFAULT '{}',
    schedule     VARCHAR(100)     NOT NULL DEFAULT '',
    status       bedrok.job_status NOT NULL DEFAULT 'active',
    max_attempts INT              NOT NULL DEFAULT 3,
    next_run_at  TIMESTAMPTZ,
    created_at   TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ      NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_jobs_status          ON bedrok.jobs (status);
CREATE INDEX IF NOT EXISTS idx_jobs_next_run_at     ON bedrok.jobs (next_run_at);
CREATE INDEX IF NOT EXISTS idx_jobs_status_next_run ON bedrok.jobs (status, next_run_at);

-- +goose Down
DROP TABLE IF EXISTS bedrok.jobs;
DROP TYPE  IF EXISTS bedrok.job_status;
