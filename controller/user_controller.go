package controller

import (
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_userapi/service"
	"github.com/saurabhjangir/bookstore_userapi/domain/users"
	"github.com/saurabhjangir/bookstore_userapi/domain/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context){
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body);
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest,errors.NewRestErrBadRequest("http bad request"))
		return
	}
	if err = json.Unmarshal(bytes,&user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, restErr := service.Create(&user)
	if restErr != nil {
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	c.JSON(http.StatusCreated, result)
	return
}

func GetUser(c *gin.Context){
	id, true := c.Params.Get("userid")
	if !true {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest("http bad request"))
		return
	}
	Id, err := strconv.ParseInt(id, 10 , 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest("http bad request"))
		return
	}
	user, Resterr := service.Get(Id)
	if err != nil{
		c.JSON(Resterr.Status, Resterr)
		return
	}
	c.JSON(http.StatusCreated, user)
	return
}