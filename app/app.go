package app

import (
	"github.com/gin-gonic/gin"
	l "github.com/saurabhjangir/bookstore_userapi/utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	l.Log.Info("Initializing bookstore user application ...")
	mapURL()
	router.Run(":3300")
}