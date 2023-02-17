// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package mydb

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
  name, password
) VALUES (
  ?, ?
)
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.Name, arg.Password)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE name = ?
`

func (q *Queries) DeleteUser(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, name)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, name, password, follow_count, follower_count FROM users
WHERE name = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Password,
		&i.FollowCount,
		&i.FollowerCount,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, name, password, follow_count, follower_count FROM users
WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Name,
		&i.Password,
		&i.FollowCount,
		&i.FollowerCount,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, name, password, follow_count, follower_count FROM users
ORDER BY user_id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Name,
			&i.Password,
			&i.FollowCount,
			&i.FollowerCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFollowCount = `-- name: UpdateFollowCount :exec
UPDATE users SET follow_count = ?
WHERE user_id = ?
`

type UpdateFollowCountParams struct {
	FollowCount sql.NullInt64 `json:"follow_count"`
	UserID      int64         `json:"user_id"`
}

func (q *Queries) UpdateFollowCount(ctx context.Context, arg UpdateFollowCountParams) error {
	_, err := q.db.ExecContext(ctx, updateFollowCount, arg.FollowCount, arg.UserID)
	return err
}

const updateFollowerCount = `-- name: UpdateFollowerCount :exec
UPDATE users SET follower_count = ?
WHERE user_id = ?
`

type UpdateFollowerCountParams struct {
	FollowerCount sql.NullInt64 `json:"follower_count"`
	UserID        int64         `json:"user_id"`
}

func (q *Queries) UpdateFollowerCount(ctx context.Context, arg UpdateFollowerCountParams) error {
	_, err := q.db.ExecContext(ctx, updateFollowerCount, arg.FollowerCount, arg.UserID)
	return err
}
