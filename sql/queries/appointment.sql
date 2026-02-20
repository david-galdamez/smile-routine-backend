-- name: CreateAppointment :one
INSERT INTO appointments (user_id, appointment_date, completed)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAppointment :one
UPDATE appointments SET completed = $1 WHERE id = $2 RETURNING *;
