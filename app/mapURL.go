package app

import (
	"github.com/saurabhjangir/bookstore_userapi/controller"
)
var (
	authenticate = router.Group("", controller.AuthenticateRequest)
)
func mapURL(){
	authenticate = router.Group("", controller.AuthenticateRequest)

	authenticate.GET("/ping", controller.Pong)

	authenticate.GET("/user/:userid", controller.GetUser)
	authenticate.GET("/user/:userid/", controller.GetUser)
	router.POST("/user", controller.CreateUser)
	router.POST("/user/", controller.CreateUser)
	router.DELETE("/user/:userid", controller.DeleteUser)
	router.DELETE("/user/:userid/", controller.DeleteUser)
	router.PUT("/user/:userid", controller.UpdateUser)
	router.PUT("/user/:userid/", controller.UpdateUser)
	router.PATCH("/user/:userid", controller.UpdateUser)
	router.PATCH("/user/:userid/", controller.UpdateUser)
	router.GET("/internal/user/search", controller.SearchUser)
	router.POST("/user/login/", controller.LoginUser)
	router.POST("/user/login", controller.LoginUser)
}
