-- +goose Up

CREATE TABLE habits (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    habit_date DATE NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    meal_id INTEGER NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (meal_id) REFERENCES meals(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE habits;
