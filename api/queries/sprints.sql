-- name: GetSprint :one
SELECT * FROM sprints
WHERE id = $1 LIMIT 1;

-- name: ListSprints :many
SELECT * FROM sprints;

-- name: CreateSprint :one
INSERT INTO sprints (
  title, start_date, end_date
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: UpdateSprint :one
UPDATE sprints SET
  title = $2,
  start_date = $3,
  end_date = $4,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteSprint :exec
DELETE FROM sprints
WHERE id = $1;
