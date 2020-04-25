package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/saurabhjangir/bookstore_userapi/utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	log.Log.Info("Initializing application")
	mapURL()
	router.Run(":3300")
}