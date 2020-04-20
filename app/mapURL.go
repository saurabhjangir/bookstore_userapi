package app

import "github.com/saurabhjangir/bookstore_userapi/controller"

func mapURL(){
	router.GET("/ping", controller.Pong)
	router.POST("/user", controller.CreateUser)
	router.POST("/user/", controller.CreateUser)
	router.GET("/user/:userid", controller.GetUser)
	router.GET("/user/:userid/", controller.GetUser)
}
