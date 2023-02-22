// Code generated by sqlc. DO NOT EDIT.

package mydb

import (
	"database/sql"
)

type Favorite struct {
	FavoriteID int64 `json:"favorite_id"`
	UserID     int64 `json:"user_id"`
	VideoID    int64 `json:"video_id"`
	Statement  bool  `json:"statement"`
}
type Relation struct {
	FollowID   int64 `json:"follow_id"`
	FollowedID int64 `json:"followed_id"`
	FollowerID int64 `json:"follower_id"`
	Deleted    int32 `json:"deleted"`
}

type User struct {
	UserID        int64         `json:"user_id"`
	Name          string        `json:"name"`
	Password      string        `json:"password"`
	FollowCount   sql.NullInt64 `json:"follow_count"`
	FollowerCount sql.NullInt64 `json:"follower_count"`
}

type Video struct {
	VideoID       int64         `json:"video_id"`
	Author        int64         `json:"author"`
	PlayUrl       string        `json:"play_url"`
	CoverUrl      string        `json:"cover_url"`
	FavoriteCount sql.NullInt64 `json:"favorite_count"`
	CommentCount  sql.NullInt64 `json:"comment_count"`
	Title         string        `json:"title"`
	CreatedAt     sql.NullTime  `json:"created_at"`
}

type Comment struct {
	CommentID int64        `json:"comment_id"`
	UserID    int64        `json:"user_id"`
	VideoID   int64        `json:"video_id"`
	Content   string       `json:"content"`
	CreatedAt sql.NullTime `json:"created_at"`
}
