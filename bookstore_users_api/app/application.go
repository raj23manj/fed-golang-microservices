package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StarApplication() {
	mapUrls()
	router.Run(":8080")
}
