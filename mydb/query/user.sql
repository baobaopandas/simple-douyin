-- name: GetUser :one
SELECT * FROM users
WHERE name = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE user_id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO users (
  name, password
) VALUES (
  ?, ?
);



-- name: DeleteUser :exec
DELETE FROM users
WHERE name = ?;