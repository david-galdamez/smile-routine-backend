-- +goose Up
CREATE TABLE meal_time (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE,

    breakfast TIME NOT NULL,
    lunch TIME NOT NULL,
    dinner TIME NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE meal_time;
