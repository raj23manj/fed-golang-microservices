package users

import (
	"fmt"
	"strings"

	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/datasources/mysql/users_db"
	// utils "github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/date"

	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/date"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser = ("INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);")
)

// not passing a  pointer, but passing a copy of the value from the callee
// how to structure our domain, 21:25
// func (user User) Get() *errors.RestErr {
// here we are passing pointer to the user object, and what changed here will affect in the caller function
func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	// when to use prepare & exec() directly, section 3, how to insert rows 14:16
	// if need to resuse prepare again then use prepare stmt
	// prepare is used to validate the query before inhand
	// prepare and exec has better performance
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		// Error 1062 (23000): Duplicate entry 'example@demo.com' for key 'users.email_unique'"
		if strings.Contains(err.Error(), "email_unique") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	// instead of uing above statements use below
	// result, err := users_db.Client.Exec(queryUpdateUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	user.Id = userId

	// current := usersDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }

	// usersDB[user.Id] = user
	// user.DateCreated = date.GetNowString()
	return nil
}
