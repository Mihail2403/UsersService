CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    img TEXT,
    user_id INT
);

-- migrate -path schema/ -database postgres://postgres:postgres@localhost:6543/auth?sslmode=disable up