CREATE TABLE forms (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    message VARCHAR NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    rating INT NOT NULL,
    message VARCHAR DEFAULT NULL,
    posted_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
);

INSERT INTO users
    (username, password)
VALUES
    ('ghytro', crypt('mypassword', gen_salt('bf')));
