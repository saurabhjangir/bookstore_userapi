package service

import (
	"fmt"
	"github.com/saurabhjangir/bookstore_userapi/domain/users"
	"github.com/saurabhjangir/bookstore_userapi/utils/errors"
	"strings"
	"time"
)
var (
	Userservice UsersServiceInterface = &UsersService{}
)
type UsersService struct{}

type UsersServiceInterface interface {
	Create(*users.User) (*users.User, *errors.RestErr)
	Update(bool, users.User) (*users.User, *errors.RestErr)
	Get(int64) (*users.User, *errors.RestErr)
	Delete(int64) *errors.RestErr
	Search(string) ([]users.User, *errors.RestErr)
	LoginUser(*users.User) (*users.User, *errors.RestErr)
}

func (s *UsersService)Create(input *users.User) (*users.User, *errors.RestErr){
	if err := input.Validate(); err != nil {
		return nil, errors.NewRestErrBadRequest("email address already exist")
	}
	input.Datecreated = time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Println(input.Datecreated)
	input.Status = users.StatusActive
	if err := input.Save(); err != nil {
		return nil, err
	}
	return input, nil
}

func (s *UsersService)Get(id int64) (*users.User, *errors.RestErr){
	user := &users.User{
		Id: id,
	}
	fmt.Println(user.Id)
	if err := user.Get(); err  != nil {
		return nil, err
	}
	return user, nil
}

func (s *UsersService)Delete(id int64) *errors.RestErr{
	user := &users.User{Id: id}
	return user.Delete()
}

func (s *UsersService)Update(request bool, input users.User) (*users.User, *errors.RestErr){
	current := &users.User{Id: input.Id}
	current, err := s.Get(current.Id);
	if err != nil {
		return nil, err
	}
	if request {
		if input.Firstname != "" {
			current.Firstname = strings.TrimSpace(input.Firstname)
		}
		if input.Lastname != "" {
			current.Firstname = strings.TrimSpace(input.Lastname)
		}
		if input.Email != "" {
			current.Email = strings.TrimSpace(input.Email)
		}
	}else {
		current.Firstname = strings.TrimSpace(input.Firstname)
		current.Lastname = strings.TrimSpace(input.Lastname)
		current.Email = strings.TrimSpace(input.Email)
	}
	if err :=  current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *UsersService)Search(status string) ([]users.User, *errors.RestErr){
	var user users.User
	user.Status = status
	return user.FindByStatus()
}

func (s *UsersService)LoginUser(user *users.User) (*users.User, *errors.RestErr){
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	if user.Email == "" || user.Password == "" {
		return nil, errors.NewRestErrBadRequest("email or password can't be empty")
	}
	dbErr := user.FindByEmailandPassword()
	if dbErr != nil {
		return nil , dbErr
	}
	return user, nil
}