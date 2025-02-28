-- name: GetRequirement :one
SELECT * FROM requirements
WHERE id = $1 LIMIT 1;

-- name: ListRequirements :many
SELECT * FROM requirements;

-- name: CreateRequirement :one
INSERT INTO requirements (
  user_id, sprint_id, feature_id, state, title, estimate, description
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
)
RETURNING *;

-- name: UpdateRequirement :one
UPDATE requirements SET
  user_id = $2,
  sprint_id = $3,
  feature_id = $4,
  state = $5,
  title = $6,
  estimate = $7,
  description = $8,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteRequirement :exec
DELETE FROM requirements
WHERE id = $1;
