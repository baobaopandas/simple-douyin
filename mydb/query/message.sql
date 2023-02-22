-- name: ListMessages :many
SELECT * FROM messages
WHERE ((from_user_id =? AND to_user_id =?) OR (from_user_id =? AND to_user_id =?)) AND create_time > ? ORDER BY create_time;

-- name: CreateMessage :execresult
INSERT INTO messages(
    from_user_id,to_user_id,content,create_time
)VALUES (?,?,?,?);