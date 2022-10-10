package app

import (
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/controllers/ping"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/controllers/users"
)

func mapUrls() {
	// ping controller
	router.GET("/ping", ping.Ping)

	// users controller
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
