package users

import "encoding/json"

type PublicUser struct {
	// Id          int64  `json:"user_id"`
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	// method 1
	// this method is applicable only when all are same or different json attributes
	// 19:05, how to marshall structs
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	// method 2
	// this method is applicable only when all the json fields are same in public & private
	// the json attributes should match that of `User` for below method
	/*
		type User struct {
			Id          int64  `json:"id"`
			FirstName   string `json:"first_name"`
			LastName    string `json:"last_name"`
			Email       string `json:"email"`
			DateCreated string `json:"date_created"`
			Status      string `json:"status"`
			Password    string `json:"password"`
		}
	*/
	// 19:05, how to marshall structs
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
