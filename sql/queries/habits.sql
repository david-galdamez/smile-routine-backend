-- name: RegisterHabit :exec
INSERT INTO habits (user_id, habit_date, completed, meal_id)
VALUES ($1, $2, $3, $4);

-- name: GetHabits :many
WITH days AS (
    SELECT generate_series($2::date, ($3::date - interval '1 day'), interval '1 day')::date AS day
)
SELECT
    d.day::date as day,
    COALESCE(COUNT(h.id) FILTER (WHERE h.completed = true), 0)::bigint AS completed_count
FROM days d
LEFT JOIN habits h
    ON h.habit_date::date = d.day
    AND h.user_id = $1
GROUP BY d.day
ORDER BY d.day;
