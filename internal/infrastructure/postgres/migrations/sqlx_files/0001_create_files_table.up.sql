CREATE TABLE IF NOT EXISTS files
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    type       VARCHAR(50)  NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    data       BYTEA
);
