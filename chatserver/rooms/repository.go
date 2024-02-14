package rooms

import (
	"context"
	"errors"
	"fmt"

	"github.com/dolencd/go-playground/chatserver/messages"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Room struct {
	Id   string `json:"id" pgx:"id"`
	Name string `json:"name" pgx:"name" binding:"required"`
}

type RoomRepo struct {
	conn *pgx.Conn
}

func NewRoomRepo(conn *pgx.Conn) RoomRepo {
	return RoomRepo{conn: conn}
}

func (r *RoomRepo) CreateRoom(room Room) (Room, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return Room{}, errors.New("failed to generate new room id")
	}
	room.Id = newId.String()
	_, err = r.conn.Exec(context.Background(), "INSERT INTO room (id, name) VALUES ($1, $2)", room.Id, room.Name)
	if err != nil {
		return Room{}, err
	}
	return room, nil
}

func (r *RoomRepo) GetRooms() ([]Room, error) {
	rows, err := r.conn.Query(context.Background(), "SELECT id, name FROM room")
	if err != nil {
		return nil, err
	}
	rooms := make([]Room, 0, 3)
	for rows.Next() {
		room := Room{}
		err := rows.Scan(&room.Id, &room.Name)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)

	}

	return rooms, nil
}

func (r *RoomRepo) GetRoom(id string) (Room, bool) {
	row := r.conn.QueryRow(context.Background(), "SELECT id, name FROM room WHERE id=$1", id)
	room := Room{}
	err := row.Scan(&room.Id, &room.Name)
	if err != nil {
		fmt.Errorf("err: %v\n", err)
		return Room{}, false
	}
	return room, true
}

func (r *RoomRepo) UpdateRoom(id string, room Room) (Room, error) {
	_, err := r.conn.Exec(context.Background(), "UPDATE room SET name=$2 WHERE id=$1", id, room.Name)
	if err != nil {
		return Room{}, err
	}
	room.Id = id
	return room, nil

}

func (r *RoomRepo) DeleteRoom(id string) error {
	_, err := r.conn.Exec(context.Background(), "DELETE FROM room WHERE id=$1", id)
	return err
}

func (r *RoomRepo) AddUserToRoom(userId, roomId string) error {
	newId, err := uuid.NewV7()
	if err != nil {
		return errors.New("failed to generate new room id")
	}
	_, err = r.conn.Exec(context.Background(), `INSERT INTO user_room (id, user_id, room_id) VALUES ($1, $2, $3)`, newId, userId, roomId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepo) RemoveUserFromRoom(userId, roomId string) error {
	_, err := r.conn.Exec(context.Background(), `DELETE FROM user_room WHERE user_id = $1 AND room_id = $2`, userId, roomId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepo) GetRoomMessages(roomId string) ([]messages.Message, error) {
	rows, err := r.conn.Query(context.Background(), `SELECT id, text, room_id, created_at, sender_user_id FROM message WHERE room_id = $1 ORDER BY created_at DESC`, roomId)
	if err != nil {
		return nil, err
	}
	msgs := make([]messages.Message, 0, 3)
	for rows.Next() {
		msg := messages.Message{}
		err := rows.Scan(&msg.Id, &msg.Text, &msg.RoomId, &msg.CreatedAt, &msg.SenderUserId)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}
