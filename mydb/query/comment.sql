-- name: GetComment :one
SELECT * FROM comments
WHERE  comment_id = ? LIMIT 1;

-- name: GetCommentsById :many
SELECT * FROM comments
WHERE  video_id = ?;

# -- name: ListComments :many
# SELECT * FROM comments
# WHERE created_at <= ?
# ORDER BY created_at LIMIT 30;

-- name: CreateComment :execresult
INSERT INTO comments (
  user_id, video_id, content
) VALUES (
  ?, ?, ?
);



-- name: DeleteComment :exec
DELETE FROM comments
WHERE comment_id = ?;