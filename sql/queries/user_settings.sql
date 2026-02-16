-- name: RegisterUserSetting :exec
INSERT INTO user_settings (user_id, wait_minutes)
VALUES ($1, $2);

-- name: UpdateUserSetting :one
UPDATE user_settings
SET wait_minutes = $2
WHERE user_id = $1 RETURNING wait_minutes;

-- name: GetUserSetting :one
SELECT wait_minutes
FROM user_settings
WHERE user_id = $1;
