-- name: GetUser :one
SELECT id, email, first_name, last_name, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserWithPassword :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserWithPasswordByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, email, first_name, last_name, created_at, updated_at FROM users;

-- name: CreateUser :one
INSERT INTO users (
  email,
  first_name,
  last_name,
  password
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING id, email, first_name, last_name, created_at, updated_at;

-- name: UpdateUser :one
UPDATE users set
  email = $2,
  first_name = $3,
  last_name = $4,
  password = $5,
  updated_at = now()
WHERE id = $1
RETURNING id, email, first_name, last_name, updated_at, created_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
