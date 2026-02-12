-- name: RegisterUserSetting :exec
INSERT INTO user_settings (user_id, wait_minutes)
VALUES ($1, $2);

-- name: UpdateUserSetting :exec
UPDATE user_settings
SET wait_minutes = $2
WHERE user_id = $1;
