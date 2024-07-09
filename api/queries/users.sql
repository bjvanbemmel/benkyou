-- name: GetUser :one
SELECT id, email, username, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserWithPassword :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserWithPasswordByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, email, username, created_at, updated_at FROM users;

-- name: CreateUser :one
INSERT INTO users (
  email,
  username,
  password
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users set
  email = $2,
  username = $3,
  password = $4,
  updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
