// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const getMessage = `-- name: GetMessage :one
SELECT
    "id", "room_id", "message", "reaction_count", "answered"
FROM messages
WHERE
    id = $1
`

func (q *Queries) GetMessage(ctx context.Context, id uuid.UUID) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Message,
		&i.ReactionCount,
		&i.Answered,
	)
	return i, err
}

const getRoom = `-- name: GetRoom :one
SELECT
    "id", "theme"
FROM rooms
WHERE id = $1
`

func (q *Queries) GetRoom(ctx context.Context, id uuid.UUID) (Room, error) {
	row := q.db.QueryRow(ctx, getRoom, id)
	var i Room
	err := row.Scan(&i.ID, &i.Theme)
	return i, err
}

const getRoomMessages = `-- name: GetRoomMessages :many
SELECT
    "id", "room_id", "message", "reaction_count", "answered"
FROM messages
WHERE
    room_id = $1
`

func (q *Queries) GetRoomMessages(ctx context.Context, roomID uuid.UUID) ([]Message, error) {
	rows, err := q.db.Query(ctx, getRoomMessages, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.Message,
			&i.ReactionCount,
			&i.Answered,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRooms = `-- name: GetRooms :many
SELECT
    "id", "theme"
FROM rooms
`

func (q *Queries) GetRooms(ctx context.Context) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(&i.ID, &i.Theme); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertMessage = `-- name: InsertMessage :one
INSERT INTO messages
( "room_id", "message" ) VALUES
    ( $1, $2 )
    RETURNING "id"
`

type InsertMessageParams struct {
	RoomID  uuid.UUID `db:"room_id" json:"room_id"`
	Message string    `db:"message" json:"message"`
}

func (q *Queries) InsertMessage(ctx context.Context, arg InsertMessageParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertMessage, arg.RoomID, arg.Message)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const insertRoom = `-- name: InsertRoom :one
INSERT INTO rooms
( "theme" ) VALUES
    ( $1 )
    RETURNING "id"
`

func (q *Queries) InsertRoom(ctx context.Context, theme string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, insertRoom, theme)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const markMessageAsAnswered = `-- name: MarkMessageAsAnswered :exec
UPDATE messages
SET
    answered = true
WHERE
    id = $1
`

func (q *Queries) MarkMessageAsAnswered(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, markMessageAsAnswered, id)
	return err
}

const reactToMessage = `-- name: ReactToMessage :one
UPDATE messages
SET
    reaction_count = reaction_count + 1
WHERE
    id = $1
    RETURNING reaction_count
`

func (q *Queries) ReactToMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, reactToMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}

const removeReactionFromMessage = `-- name: RemoveReactionFromMessage :one
UPDATE messages
SET
    reaction_count = reaction_count - 1
WHERE
    id = $1
    RETURNING reaction_count
`

func (q *Queries) RemoveReactionFromMessage(ctx context.Context, id uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, removeReactionFromMessage, id)
	var reaction_count int64
	err := row.Scan(&reaction_count)
	return reaction_count, err
}
