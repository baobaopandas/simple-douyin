-- name: GetUser :one
SELECT * FROM users
WHERE name = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE user_id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id;

-- name: CreateUser :execresult
INSERT INTO users (
  name, password
) VALUES (
  ?, ?
);

-- name: UpdateFollowCount :exec
UPDATE users SET follow_count = ?
WHERE user_id = ?;

-- name: UpdateFollowerCount :exec
UPDATE users SET follower_count = ?
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE name = ?;