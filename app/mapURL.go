package app

import "github.com/saurabhjangir/bookstore_userapi/controller"

func mapURL(){
	router.GET("/ping", controller.Pong)
	router.POST("/user", controller.CreateUser)
	router.POST("/user/", controller.CreateUser)
	router.GET("/user/:userid", controller.GetUser)
	router.GET("/user/:userid/", controller.GetUser)
	router.DELETE("/user/:userid", controller.DeleteUser)
	router.DELETE("/user/:userid/", controller.DeleteUser)
	router.PUT("/user/:userid", controller.UpdateUser)
	router.PUT("/user/:userid/", controller.UpdateUser)
	router.PATCH("/user/:userid", controller.UpdateUser)
	router.PATCH("/user/:userid/", controller.UpdateUser)
	router.GET("/internal/user/search", controller.SearchUser)
	/*
	put post and search with query param
	 */
}
