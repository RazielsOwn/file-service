CREATE TABLE IF NOT EXISTS files(
    id serial PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    path VARCHAR(255)
);