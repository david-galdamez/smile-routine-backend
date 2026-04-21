-- name: CreateAppointment :one
INSERT INTO appointments (user_id, appointment_date, appointment_time, title, note)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateAppointment :one
UPDATE appointments SET completed = $1 WHERE id = $2 RETURNING *;

-- name: GetAppointments :many
SELECT * FROM appointments WHERE user_id = $1 AND completed = FALSE;

-- name: GetAppointmentById :one
SELECT * FROM appointments WHERE id = $1;
