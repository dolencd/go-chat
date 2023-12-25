package users

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRepo struct {
	conn *pgx.Conn
}

var users = make(map[string]User)

func (r *UserRepo) Close() {
	r.conn.Close(context.Background())
}

func NewUserRepo() (UserRepo, error) {
	pg_url := os.Getenv("POSTGRES_URL")
	conn, err := pgx.Connect(context.Background(), pg_url)
	if err != nil {
		return UserRepo{}, err
	}

	return UserRepo{conn: conn}, nil
}

func (r *UserRepo) CreateUser(user User) (User, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return User{}, errors.New("failed to generate new user id")
	}
	user.Id = newId.String()
	users[user.Id] = user
	return user, nil
}

func (r *UserRepo) GetUsers() []User {
	values := make([]User, 0, len(users))

	for _, v := range users {
		values = append(values, v)
	}

	return values
}

func (r *UserRepo) GetUser(id string) (User, bool) {
	user, isFound := users[id]
	return user, isFound
}

func (r *UserRepo) UpdateUser(id string, user User) (User, error) {
	user.Id = id
	_, ok := users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	users[id] = user
	return user, nil

}

func (r *UserRepo) DeleteUser(id string) error {
	_, ok := users[id]
	if !ok {
		return errors.New("user not found")
	}

	delete(users, id)
	return nil
}
