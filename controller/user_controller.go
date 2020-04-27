package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_userapi/domain/users"
	"github.com/saurabhjangir/bookstore_userapi/service"
	log "github.com/saurabhjangir/bookstore_userapi/utils/logger"
	"github.com/saurabhjangir/oauth-client"
	"github.com/saurabhjangir/utils-lib-golang/errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	AuthClient oauth_client.IoauthClient = &oauth_client.OauthClient{}
)

func getUserInputId(userId string) (int64, *errors.RestErr) {
	Id, err := strconv.ParseInt(userId, 10 , 64)
	if err != nil {
		return Id, errors.NewRestErrBadRequest(err.Error())
	}
	return Id, nil
}

func AuthenticateRequest(c *gin.Context){
	if err := AuthClient.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		c.Abort()
		return
	}
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
	Id, err := strconv.ParseInt(c.Param("userid"), 10 , 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewRestErrBadRequest("http bad request"))
		return
	}
	user, Resterr := service.Userservice.Get(Id)
	if Resterr != nil{
		c.JSON(Resterr.Status, Resterr)
		return
	}
	CallerId, _ := AuthClient.GetCallerID(c.Request)
	if *CallerId != Id {
		c.JSON(http.StatusCreated, user.Marshal(false))
		return
	}
	c.JSON(http.StatusCreated, user.Marshal(true))
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