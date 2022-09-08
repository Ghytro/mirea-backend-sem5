CREATE TABLE some_table (
    id SERIAL PRIMARY KEY,
    column1 VARCHAR NOT NULL,
    column2 INT NOT NULL
);

INSERT INTO some_table
    (column1, column2)
VALUES
    ('Welcome to postgres!', 10);
