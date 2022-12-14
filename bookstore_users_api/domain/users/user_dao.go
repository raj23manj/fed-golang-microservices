package users

import (
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/datasources/mysql/users_db"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/logger"

	// utils "github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/date"

	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/errors"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/mysql_utils"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?) ;"
	indexUniqueEmail      = "email_unique"
	queryGetter           = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE id = ? ;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id = ? ;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ? ;"
	queryFindUserByStatus = "SELECT first_name, last_name, email, date_created, status, id FROM users WHERE status = ? ;"
)

// not passing a pointer, but passing a copy of the value from the callee
// how to structure our domain, 21:25
// func (user User) Get() *errors.RestErr {
// here we are passing pointer to the user object, and what changed here will affect in the caller function
func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetter)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	// using Query => https://github.com/golang/go/wiki/SQLInterface
	// results, err := stmt.Query(user.Id)
	// if err != nil {
	// return errors.NewInternalServerError(err.Error())
	// }
	// defer results.Close()

	result := stmt.QueryRow(user.Id)
	// scan populates the attributes matched from the query and adds them to user object
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		// fmt.Println(err)
		// // in this case we do not check for errors like we did for save, here we check for string, 20:10 how to handle sql errors
		// if strings.Contains(err.Error(), "no rows in result set") {
		// 	return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error while trying to get user %d : %s", user.Id, err.Error()))
		logger.Error("error when trying to get user by id", err)
		return mysql_utils.ParseError(err)
	}

	// result := usersDB[user.Id]
	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	// }
	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	// when to use prepare & exec() directly, section 3, how to insert rows 14:16
	// if need to resuse prepare again then use prepare stmt
	// prepare is used to validate the query before inhand
	// prepare and exec has better performance
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	// user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		// // identifying errors
		// // http://go-database-sql.org/errors.html

		// // to check if it is a mysql error, 3:48 how to handle sql errors
		// // convert err to *mysql.MySQLError
		// // if err is a type of *mysql.MySQLError then ok will be true
		// sqlErr, ok := err.(*mysql.MySQLError)
		// // check if the error is not a type of mysql error, 5:40 how to handle sql errors
		// if !ok {
		// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
		// }
		// fmt.Println("################################################################")
		// fmt.Println(sqlErr)
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
		// /*
		// 	  14:21 how to handle sql errors
		// 		we can use switch case to pin point to specific errors
		// 		switch sqlErr.Number {
		// 		case 1062:
		// 			return  errors.NewInternalServerError(fmt.Sprintf("email %s already exisits", err.Error()))
		// 		}
		// */
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))

		// // other way of doing using the message as regex
		// // // Error 1062 (23000): Duplicate entry 'example@demo.com' for key 'users.email_unique'"
		// // if strings.Contains(err.Error(), indexUniqueEmail) {
		// // 	return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		// // }
		// // return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))

		// use util method
		logger.Error("error when trying to save user", err)
		// return mysql_utils.ParseError(err)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating new user", err)
		return errors.NewInternalServerError("database error")
		// return mysql_utils.ParseError(err)
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		// return errors.NewInternalServerError(err.Error())
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	// here we dont care about the result but check for the error, 14:51 how to update rows
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying update user", err)
		return errors.NewInternalServerError("database error")
		// return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		// return mysql_utils.ParseError(err)
		logger.Error("error when trying delete user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare findbystatus user statement", err)
		return nil, errors.NewInternalServerError("database error")
		// return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying find user by status", err)
		return nil, errors.NewInternalServerError("database error")
		// return nil, mysql_utils.ParseError(err)
	}
	// 9:10, how to find rows
	// this defer statement is put after error, because if error occurs then rows will be nil and cause nil panic during runtime, so can't put immediately
	// after like this
	// rows, err := stmt.Query(status)
	// defer rows.Close()
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		// passing user.FirstName, user.LastName, here and reaching append code the user will be nil since we are passing a copy of user
		// hence pass &user.FirstName, &user.LastName ...
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Id); err != nil {
			logger.Error("error when trying to scan find user by status", err)
			return nil, errors.NewInternalServerError("database error")
			// return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("Users not found with matching status")
	}

	return results, nil
}
