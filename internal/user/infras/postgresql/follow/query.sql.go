// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package postgresql

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const allByFollowedUserID = `-- name: AllByFollowedUserID :many
select
    id,
    followed_user_id,
    following_user_id,
    status,
    created_date,
    updated_date
from
    follow
where
    followed_user_id = $1
    and status = ANY ($2::text[])
    and id < $3
order by
    id desc
limit
    $4
`

type AllByFollowedUserIDParams struct {
	FollowedUserID uuid.UUID `json:"followed_user_id"`
	Column2        []string  `json:"column_2"`
	ID             int64     `json:"id"`
	Limit          int32     `json:"limit"`
}

func (q *Queries) AllByFollowedUserID(ctx context.Context, arg AllByFollowedUserIDParams) ([]Follow, error) {
	rows, err := q.db.QueryContext(ctx, allByFollowedUserID,
		arg.FollowedUserID,
		pq.Array(arg.Column2),
		arg.ID,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follow
	for rows.Next() {
		var i Follow
		if err := rows.Scan(
			&i.ID,
			&i.FollowedUserID,
			&i.FollowingUserID,
			&i.Status,
			&i.CreatedDate,
			&i.UpdatedDate,
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

const countByFollowedUserID = `-- name: CountByFollowedUserID :one
select
    count(*)
from
    follow
where
    followed_user_id = $1
`

func (q *Queries) CountByFollowedUserID(ctx context.Context, followedUserID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, countByFollowedUserID, followedUserID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createFollow = `-- name: CreateFollow :one
insert into
    follow (followed_user_id, following_user_id)
values
    ($1, $2)
returning
    id, followed_user_id, following_user_id, status, created_date, updated_date
`

type CreateFollowParams struct {
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
	FollowingUserID uuid.UUID `json:"following_user_id"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error) {
	row := q.db.QueryRowContext(ctx, createFollow, arg.FollowedUserID, arg.FollowingUserID)
	var i Follow
	err := row.Scan(
		&i.ID,
		&i.FollowedUserID,
		&i.FollowingUserID,
		&i.Status,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const getOne = `-- name: GetOne :one
select
    id, followed_user_id, following_user_id, status, created_date, updated_date
from
    follow
where
    followed_user_id = $1
    and following_user_id = $2
`

type GetOneParams struct {
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
	FollowingUserID uuid.UUID `json:"following_user_id"`
}

func (q *Queries) GetOne(ctx context.Context, arg GetOneParams) (Follow, error) {
	row := q.db.QueryRowContext(ctx, getOne, arg.FollowedUserID, arg.FollowingUserID)
	var i Follow
	err := row.Scan(
		&i.ID,
		&i.FollowedUserID,
		&i.FollowingUserID,
		&i.Status,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}

const updateFollow = `-- name: UpdateFollow :one
update follow
set
    status = coalesce($1, "follow".status)
where
    followed_user_id = $2
    and following_user_id = $3
returning
    id, followed_user_id, following_user_id, status, created_date, updated_date
`

type UpdateFollowParams struct {
	Status          string    `json:"status"`
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
	FollowingUserID uuid.UUID `json:"following_user_id"`
}

func (q *Queries) UpdateFollow(ctx context.Context, arg UpdateFollowParams) (Follow, error) {
	row := q.db.QueryRowContext(ctx, updateFollow, arg.Status, arg.FollowedUserID, arg.FollowingUserID)
	var i Follow
	err := row.Scan(
		&i.ID,
		&i.FollowedUserID,
		&i.FollowingUserID,
		&i.Status,
		&i.CreatedDate,
		&i.UpdatedDate,
	)
	return i, err
}
