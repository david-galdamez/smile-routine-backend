-- +goose Up
CREATE TABLE appointments (
    id BIGSERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    appointment_date DATE NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE appointments;
