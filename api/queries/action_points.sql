-- name: GetActionPoint :one
SELECT * FROM action_points
WHERE id = $1;

-- name: ListActionPoints :many
SELECT * FROM action_points;

-- name: CreateActionPoint :one
INSERT INTO action_points (
  title
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateActionPoint :one
UPDATE action_points SET
  title = $2,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteActionPoint :exec
DELETE FROM action_points
WHERE id = $1;
