-- name: GetFollowedIdByFollower :many
SELECT followed_id FROM relations
WHERE  follower_id = ?
AND deleted = 0;

-- name: GetFollowerIdByFollowed :many
SELECT follower_id FROM relations
WHERE  followed_id = ?
AND deleted = 0;

-- name: GetFollowedCount :one
SELECT count(*) FROM relations
WHERE follower_id = ?
AND deleted = 0;

-- name: GetFollowerCount :one
SELECT count(*) FROM relations
WHERE followed_id = ?
AND deleted = 0;

-- name: GetRelationByID :one
SELECT deleted FROM relations
WHERE followed_id = ?
AND follower_id = ?;

-- name: CreateRelation :exec
INSERT INTO relations (
  followed_id, follower_id
) VALUES (
  ?, ?
);

-- name: UpdateRelation :exec
UPDATE relations SET deleted = ?
WHERE followed_id = ?
AND follower_id = ?;