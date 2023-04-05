CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name char(256) NOT NULL,
    last_name char(256) DEFAULT NULL,
    login char(256) NOT NULL,
    password char(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
