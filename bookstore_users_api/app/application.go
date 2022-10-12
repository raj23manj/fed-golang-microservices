package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raj23manj/fed-golang-microservices/bookstore_users_api/logger"
)

var (
	router = gin.Default()
)

func StarApplication() {
	mapUrls()

	logger.Info("Starting application..........")
	router.Run(":8080")
}
