-- name: GetFollowedIdByFollower :many
SELECT followed_id FROM relations
WHERE  follower_id = $1
AND deleted = 0;

-- name: GetFollowerIdByFollowed :many
SELECT follower_id FROM relations
WHERE  followed_id = $1
AND deleted = 0;

-- name: GetFollowedCount :one
SELECT deleted FROM relations
WHERE follower_id = $1;

-- name: GetFollowerCount :one
SELECT count(*) FROM relations
WHERE followed_id = $1;

-- name: GetRelationByID :one
SELECT deleted FROM relations
WHERE followed_id = $1
AND follower_id = $2;

-- name: CreateRelation :exec
INSERT INTO relations (
  followed_id, follower_id
) VALUES (
  $1, $2
);

-- name: UpdateRelation :exec
UPDATE relations SET deleted = $3
WHERE followed_id = $1
AND follower_id = $2;