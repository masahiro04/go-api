-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE blogs (
    id  SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);


CREATE TABLE users (
    id  SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    UNIQUE(email)
);
