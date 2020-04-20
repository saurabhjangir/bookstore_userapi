package service

import (
	"github.com/saurabhjangir/bookstore_userapi/domain/users"
	"github.com/saurabhjangir/bookstore_userapi/domain/errors"
)

func Create(input *users.User) (*users.User, *errors.RestErr){
	if err := input.Validate(); err != nil {
		return nil, errors.NewRestErrBadRequest("email address already exist")
	}
	if err := input.Save(); err != nil {
		return nil, err
	}
	return input, nil
}

func Get(id int64) (*users.User, *errors.RestErr){
	user := &users.User{
		Id: id,
	}
	if err := user.Get(); err  != nil {
		return nil, err
	}
	return user, nil
}