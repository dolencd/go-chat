package users

import (
	"context"
	"errors"
	"fmt"
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
	_, err = r.conn.Exec(context.Background(), "INSERT INTO users (Id, Username, Email) VALUES ($1, $2, $3)", user.Id, user.Username, user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetUsers() ([]User, error) {
	rows, err := r.conn.Query(context.Background(), "SELECT Id, Username, Email FROM Users")
	if err != nil {
		return nil, err
	}
	users := make([]User, 0, 3)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}

	return users, nil
}

func (r *UserRepo) GetUser(id string) (User, bool) {
	row := r.conn.QueryRow(context.Background(), "SELECT Id, Username, Email FROM Users WHERE Id=$1", id)
	user := User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return User{}, false
	}
	return user, true
}

func (r *UserRepo) UpdateUser(id string, user User) (User, error) {
	_, err := r.conn.Exec(context.Background(), "UPDATE Users SET Username=$2, Email=$3 WHERE Id=$1", id, user.Username, user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil

}

func (r *UserRepo) DeleteUser(id string) error {
	_, err := r.conn.Exec(context.Background(), "DELETE FROM Users WHERE id=$1", id)
	return err
}
