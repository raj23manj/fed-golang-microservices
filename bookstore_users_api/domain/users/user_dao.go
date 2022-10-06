package users

import (
	"fmt"

	utils "github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/date"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// not passing a  pointer, but passing a copy of the value from the callee
// how to structure our domain, 21:25
// func (user User) Get() *errors.RestErr {
// here we are passing pointer to the user object, and what changed here will affect in the caller function
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	user.DateCreated = utils.GetNowString()
	return nil
}
