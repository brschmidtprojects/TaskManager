CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE tasks
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    size       BIGINT      NOT NULL,
    type       VARCHAR     NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
