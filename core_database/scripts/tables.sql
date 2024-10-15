-- tables.sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) UNIQUE,
                       name VARCHAR(100),
                       email VARCHAR(100) UNIQUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                       followers_count INT DEFAULT 0,
                       following_count INT DEFAULT 0,
);
