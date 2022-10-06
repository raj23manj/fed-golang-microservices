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

func CreateUser(c *gin.Context) {
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

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

	// c.String(http.StatusNotImplemented, "implement me")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me")
// }
