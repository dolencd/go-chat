-- MARK: Message

-- name: CreateMessage :one
INSERT INTO message (
  text, room_id, sender_user_id
) VALUES (?, ?, ?)
RETURNING *;

-- name: GetRoomMessages :many
SELECT id, text, room_id, created_at, sender_user_id FROM message WHERE room_id = ?;

-- name: GetMessage :one
SELECT * FROM message
WHERE id = ? LIMIT 1;

-- MARK: Room

-- name: CreateRoom :one
INSERT INTO room (
  id, name
) VALUES (?, ?)
RETURNING *;

-- name: GetRooms :many
SELECT id, name FROM room;

-- name: GetRoom :one
SELECT id, name FROM room WHERE id=? LIMIT 1;

-- name: AddUserToRoom :exec
INSERT INTO user_room (user_id, room_id) VALUES (?, ?) RETURNING *;

-- name: RemoveUserFromRoom :exec
DELETE FROM user_room WHERE user_id = ? AND room_id = ?;

-- MARK: User

-- name: GetUser :one
SELECT * FROM app_user
WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO app_user (
  username, email
) VALUES (?, ?)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM app_user;

-- name: UpdateUser :exec
UPDATE app_user
  set username = ?,
  email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM app_user
WHERE id = ?;