package messages

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Message struct {
	Id        string `json:"id" pgx:"id"`
	CreatedAt string `json:"created_at" pgx:"created_at"`
	Text      string `json:"text" pgx:"text"`
	RoomId    string `json:"room_id" pgx:"room_id"`
}

type MessageRepo struct {
	conn *pgx.Conn
}

func NewMessageRepo(conn *pgx.Conn) MessageRepo {
	return MessageRepo{conn: conn}
}

func (r *MessageRepo) CreateMessage(msg Message) (Message, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return Message{}, errors.New("failed to generate new message id")
	}
	msg.Id = newId.String()
	_, err = r.conn.Exec(context.Background(), "INSERT INTO message (id, text, room_id) VALUES ($1, $2)", msg.Id, msg.Text, msg.RoomId)
	if err != nil {
		return Message{}, err
	}
	return msg, nil
}

func (r *MessageRepo) GetMessages() ([]Message, error) {
	rows, err := r.conn.Query(context.Background(), "SELECT id, text, room_id created_at FROM messages")
	if err != nil {
		return nil, err
	}
	msgs := make([]Message, 0, 3)
	for rows.Next() {
		msg := Message{}
		err := rows.Scan(&msg.Id, &msg.Text, &msg.RoomId, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)

	}

	return msgs, nil
}

func (r *MessageRepo) GetMessage(id string) (Message, bool) {
	row := r.conn.QueryRow(context.Background(), "SELECT id, text, room_id, created_at FROM messages WHERE id=$1", id)
	msg := Message{}
	err := row.Scan(&msg.Id, &msg.Text, &msg.RoomId, &msg.CreatedAt)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return Message{}, false
	}
	return msg, true
}
