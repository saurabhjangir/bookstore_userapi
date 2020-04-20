package users

import (
	"github.com/saurabhjangir/bookstore_userapi/domain/errors"
	"strings"
	_ "github.com/saurabhjangir/bookstore_userapi/datasources/mysql/users_db"
)

type  User struct {
	Id int64 `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email"`
	Datecreated string `json:"date_created"`
}

// Validate .. Why should this method bind to datatype instead of package ?
func (input *User)Validate() *errors.RestErr {
	input.Firstname = strings.TrimSpace(input.Firstname)
	input.Lastname = strings.TrimSpace(input.Lastname)
	input.Email = strings.TrimSpace(input.Email)
	if userDB[input.Email] != nil {
		return errors.NewRestErrBadRequest("email address already exist")
	}
	return nil
}