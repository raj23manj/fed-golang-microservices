package services

import (
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/domain/users"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/crypto_utils"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/date"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/errors"
)

// how to write a  service
// 23:00, Service Structure
// first write a interface
// write a empty struct
// define methods implementing the interface
// expose a vaiable to outside services `UserService userServiceInterface = &userService{}`

// grouping of services and help in testing, Service Structure 7:54
// structurin for mocking 18:50, Service Structure
var (
	// UserService userService = userService{}
	UserService userServiceInterface = &userService{}
)

type userService struct{}

// 15:40, service structure for mocking and structuring
type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	// if err := user.Validate(); err != nil {
	// 	return nil, err
	// }

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	current, err := s.GetUser(user.Id)
	if err != nil {
		return err
	}

	return current.Delete()
}

// []users.User  ~= users.Users see in the dto
// 27:00, How to marshall structs
func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
