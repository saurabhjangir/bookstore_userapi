package users

import (
	"github.com/saurabhjangir/bookstore_userapi/datasources/mysql/users_db"
	"github.com/saurabhjangir/bookstore_userapi/domain/errors"
	"time"
)

var (
	userDB = map[string]*User{}
	queryInsertUser = "INSERT INTO users(id, first_name, last_name, date_created, email) VALUES(?,?,?,?,?);"
	queryGetUserByID = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
)

func (input *User)Save() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewRestErrInteralServer("Internal server error")
	}
	defer stmt.Close()
	input.Datecreated = time.Now().UTC().String()
	result , err := stmt.Exec(input.Id, input.Firstname, input.Lastname, input.Datecreated, input.Email)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	input.Id, err = result.LastInsertId()
	return nil
}

func (input *User)Get() *errors.RestErr{
	stmt, err := users_db.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(input.Id)
	if getErr := result.Scan(&input.Id, &input.Firstname, &input.Lastname, &input.Email, &input.Datecreated,); getErr != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	return nil
}