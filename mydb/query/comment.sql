-- name: GetComment :one
SELECT * FROM comments
WHERE  comment_id = ? LIMIT 1;

-- name: GetCommentsById :many
SELECT * FROM comments
WHERE  video_id = ?
ORDER BY created_at DESC;

-- name: CreateComment :execresult
INSERT INTO comments (
  comment_id, user_id, video_id, content
) VALUES (
  ?, ?, ?, ?
);

-- name: DeleteComment :exec
DELETE FROM comments
WHERE comment_id = ?;

-- name: MaxCommentID :one
SELECT MAX(comment_id) FROM comments