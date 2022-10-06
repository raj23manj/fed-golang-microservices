package services

import (
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/domain/users"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
