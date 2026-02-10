-- name: RegisterMealTime :one
INSERT INTO meal_time (user_id, breakfast, lunch, dinner)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetMealTimes :one
SELECT * FROM meal_time WHERE user_id = $1 LIMIT 1;

-- name: UpdateMealTime :one
UPDATE meal_time SET breakfast = $2, lunch = $3, dinner = $4 WHERE user_id = $1 RETURNING *;
