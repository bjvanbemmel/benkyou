-- name: GetFeature :one
SELECT * FROM features
WHERE id = $1 LIMIT 1;

-- name: ListFeatures :many
SELECT * FROM features;

-- name: CreateFeature :one
INSERT INTO features (
  user_id, sprint_id, state, title, description, position
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  (
    SELECT CASE COUNT(*) WHEN 0 THEN 0 ELSE MAX(features.position) + 1 END AS position
		FROM features
		ORDER BY position DESC
		LIMIT 1
  )
)
RETURNING *;

-- name: UpdateFeature :one
UPDATE features SET
  user_id = $2,
  sprint_id = $3,
  state = $4,
  title = $5,
  description = $6,
  position = $7,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: IncrementPositionForAllFeaturesBetweenDescending :exec
UPDATE features f SET
  position = f.position + 1
FROM (
  SELECT * FROM features f
  WHERE f.position >= $1 AND f.position <= $2
  ORDER BY position
) fd
WHERE fd.id = f.id;

-- name: IncrementPositionForAllFeaturesBetweenAscending :exec
UPDATE features f SET
  position = f.position - 1
FROM (
  SELECT * FROM features f
  WHERE f.position >= $1 AND f.position <= $2
  ORDER BY f.position
) fd
WHERE fd.id = f.id;

-- name: DeleteFeature :exec
DELETE FROM features
WHERE id = $1;
