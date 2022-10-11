package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/domain/users"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/services"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/utils/errors"
)

func Create(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// different ways to read from context, below method 1
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: handle error
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(bytes))
	// // {
	// //   "id": 23,
	// //   "first_name": "demo",
	// //   "last_name": "demo2",
	// //   "email": "example@demo.com",
	// //   "date_created": "01/01/2022"
	// // }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO: handle error
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(user)
	// // o/p after unmarshalling
	// // {23 demo demo2 example@demo.com 01/01/2022}

	// method 2
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		fmt.Println(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))

	// c.String(http.StatusNotImplemented, "implement me")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }

func Update(c *gin.Context) {
	var user users.User

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		fmt.Println(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	// // method 1
	// // this approach is ok when the end point is mapped to only one function,
	// // but this same method is mapped to multiple api end points, it will be a problem
	// // 24:00, How to marshall structs
	// result := make([]interface{}, len(users))
	// for index, user := range users {
	// 	result[index] = user.Marshall(c.GetHeader("X-Public") == "true")
	// }

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

// shared methods

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}
