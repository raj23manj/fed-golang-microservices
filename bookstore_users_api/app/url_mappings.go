package app

import (
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/controllers"
)

func mapUrls() {
	// Ping Controller
	router.GET("/ping", controllers.Ping)
}
