package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_userapi/domain/users"
	"github.com/saurabhjangir/bookstore_userapi/service"
	"github.com/saurabhjangir/bookstore_userapi/utils/errors"
	log "github.com/saurabhjangir/bookstore_userapi/utils/logger"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getUserInputId(userId string) (int64, *errors.RestErr) {
	Id, err := strconv.ParseInt(userId, 10 , 64)
	if err != nil {
		return Id, errors.NewRestErrBadRequest(err.Error())
	}
	return Id, nil
}

func CreateUser(c *gin.Context){
	log.Log.Info("Request received: Create User")
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body);
	if err != nil {
		c.JSON(http.StatusBadRequest,errors.NewRestErrBadRequest(err.Error()))
		return
	}
	if err = json.Unmarshal(bytes,&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	result, restErr := service.Userservice.Create(&user)
	if restErr != nil {
		c.JSON(http.StatusBadRequest, restErr)
		return
	}
	c.JSON(http.StatusCreated, result)
	return
}

func GetUser(c *gin.Context){
	log.Log.Info("Request received: Get User")
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
	user, Resterr := service.Userservice.Get(Id)
	if Resterr != nil{
		c.JSON(Resterr.Status, Resterr)
		return
	}
	c.JSON(http.StatusCreated, user)
	return
}

func DeleteUser(c *gin.Context) {
	Id, err := getUserInputId(c.Param("userid"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	RestErr := service.Userservice.Delete(Id)
	if RestErr != nil {
		c.JSON(RestErr.Status, RestErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
	// In case of http.StatusNoContent , status: deleted won't be shown on user end
	return
}

func UpdateUser(c *gin.Context){
	Id, idErr := getUserInputId(c.Param("userid"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user := &users.User{Id: Id}
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest(err.Error()))
		return
	}
	if err := json.Unmarshal(bytes, user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest(err.Error()))
		return
	}
	user, RestErr := service.Userservice.Update(c.Request.Method == http.MethodPatch, *user)
	if RestErr != nil {
		c.JSON(RestErr.Status, RestErr)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func SearchUser(c *gin.Context){
	log.Log.Info("Request Received: Search user")
	log.Log.Info(c.Request)
	status := c.Query("status")
	users, err := service.Userservice.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

func LoginUser(c *gin.Context){
	log.Log.Info("login request received", c.Request)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var user users.User
	if err := json.Unmarshal(bytes, &user); err != nil {
		log.Log.Info(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user1, loginErr := service.Userservice.LoginUser(&user)
	if loginErr != nil {
		log.Log.Info(err.Error())
		c.JSON(loginErr.Status, loginErr)
		return
	}
	log.Log.Info("login Successful: ", user1)
	c.JSON(http.StatusOK, user1)
	return
}