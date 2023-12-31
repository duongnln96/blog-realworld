// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package postgresql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Follow struct {
	ID              int64        `json:"id"`
	FollowedUserID  uuid.UUID    `json:"followed_user_id"`
	FollowingUserID uuid.UUID    `json:"following_user_id"`
	Status          string       `json:"status"`
	CreatedDate     time.Time    `json:"created_date"`
	UpdatedDate     sql.NullTime `json:"updated_date"`
}
