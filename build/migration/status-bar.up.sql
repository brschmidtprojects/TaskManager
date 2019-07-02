CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE status_bars
(
    id               uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id          uuid REFERENCES tasks (id) NOT NULL,
    size             BIGINT                     NOT NULL,
    current_progress BIGINT           DEFAULT 0,
    state            VARCHAR          DEFAULT 'ready',
    done             BOOLEAN          DEFAULT false,
    created_at       TIMESTAMPTZ                NOT NULL,
    updated_at       TIMESTAMPTZ                NOT NULL,
    UNIQUE (task_id)
);
