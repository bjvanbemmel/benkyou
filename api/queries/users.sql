-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

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
  password = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
