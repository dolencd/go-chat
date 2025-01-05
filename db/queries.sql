-- name: GetUser :one
SELECT * FROM app_user
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO app_user (
  id, username, email
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM app_user;

-- name: UpdateUser :exec
UPDATE app_user
  set username = $2,
  email = $3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM app_user
WHERE id = $1;