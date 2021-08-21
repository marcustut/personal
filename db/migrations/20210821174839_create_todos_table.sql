-- migrate:up
CREATE TABLE todos (
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT, 
    completed   BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP DEFAULT NOW()
);

-- migrate:down
DROP TABLE todos;
