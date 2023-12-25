package users

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = make(map[string]User)

func CreateUser(user User) (User, error) {
	newId, err := uuid.NewV7()
	if err != nil {
		return User{}, errors.New("failed to generate new user id")
	}
	user.Id = newId.String()
	users[user.Id] = user
	return user, nil
}

func GetUsers() []User {
	values := make([]User, 0, len(users))

	for _, v := range users {
		values = append(values, v)
	}

	return values
}

func GetUser(id string) (User, bool) {
	user, isFound := users[id]
	return user, isFound
}

func UpdateUser(id string, user User) (User, error) {
	user.Id = id
	_, ok := users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	users[id] = user
	return user, nil

}

func DeleteUser(id string) error {
	_, ok := users[id]
	if !ok {
		return errors.New("user not found")
	}

	delete(users, id)
	return nil
}
