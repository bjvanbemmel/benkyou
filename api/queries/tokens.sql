-- name: GetToken :one
SELECT * FROM tokens
WHERE id = $1 LIMIT 1;

-- name: ListTokens :many
SELECT * FROM tokens;

-- name: CreateToken :one
INSERT INTO tokens (
  user_id,
  value,
  expires_at
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: UpdateToken :one
UPDATE tokens set
  user_id = $2,
  value = $3,
  expires_at = $4,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1;
