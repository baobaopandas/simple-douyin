-- name: GetVideo :one
SELECT * FROM videos
WHERE  video_id = ? LIMIT 1;

-- name: GetVideoById :many
SELECT * FROM videos
WHERE  author = ?;


-- name: ListVideos :many
SELECT * FROM videos
WHERE created_at <= ?
ORDER BY created_at LIMIT 30;




-- name: CreateVideo :execresult
INSERT INTO videos (
  author, play_url, cover_url, title
) VALUES (
  ?, ?, ?, ?
);



-- name: DeleteVideo :exec
DELETE FROM videos
WHERE video_id = ?;