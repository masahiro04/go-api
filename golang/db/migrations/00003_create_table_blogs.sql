-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE blogs (
    id  SERIAL NOT NULL,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);
