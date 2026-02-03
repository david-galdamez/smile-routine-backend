-- +goose Up

CREATE TABLE meals (
    id SERIAL PRIMARY KEY,
    meal_name VARCHAR(15) NOT NULL
);

INSERT INTO meals (meal_name) VALUES ('breakfast'), ('lunch'), ('dinner');

-- +goose Down
DROP TABLE meals;
