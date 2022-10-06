package services

import "github.com/raj23manj/fed-golang-microservices/bookstore_users_api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
