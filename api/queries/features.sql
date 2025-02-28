-- name: GetFeature :one
SELECT * FROM features
WHERE id = $1 LIMIT 1;

-- name: ListFeatures :many
SELECT * FROM features;

-- name: CreateFeature :one
INSERT INTO features (
  user_id, sprint_id, state, title, description
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: UpdateFeature :one
UPDATE features SET
  user_id = $2,
  sprint_id = $3,
  state = $4,
  title = $5,
  description = $6,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteFeature :exec
DELETE FROM features
WHERE id = $1;
