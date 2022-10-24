CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    name VARCHAR,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE forms (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    message VARCHAR NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    rating INT NOT NULL,
    message VARCHAR DEFAULT NULL,
    posted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO users
    (username, password, email, is_admin)
VALUES
    ('ghytro', crypt('root', gen_salt('bf')), 'some_email@gmail.com', TRUE);
