package app

import (
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/controllers/ping"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/controllers/users"
)

func mapUrls() {
	// ping controller
	router.GET("/ping", ping.Ping)

	// users controller
	router.GET("/users/:user_id", users.Get)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
