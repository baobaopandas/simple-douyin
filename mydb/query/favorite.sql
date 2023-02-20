-- name: AddFavorite :exec
INSERT INTO favorite (
    user_id, video_id, statement
) VALUES (
    ?, ?, 1
);

-- name: GetInfo :one
SELECT * FROM favorite 
WHERE user_id = ? AND video_id = ?;

-- name: UpdateFavorite :exec
UPDATE favorite SET statement = 1
WHERE user_id = ? AND video_id = ?;

-- name: DeleteFavorite :exec
UPDATE favorite SET statement = 0
WHERE user_id = ? AND video_id = ?;

-- name: GetUserLike :many
SELECT * FROM favorite
WHERE user_id = ?;

-- name: AddFavoriteCount :exec
UPDATE videos SET favorite_count=favorite_count+1
WHERE video_id = ?;

-- name: DelFavoriteCount :exec
UPDATE videos SET favorite_count=favorite_count-1
WHERE video_id = ?;
