package users

import (
	"github.com/saurabhjangir/bookstore_userapi/datasources/mysql/users_db"
	"github.com/saurabhjangir/bookstore_userapi/domain/errors"
	log "github.com/saurabhjangir/bookstore_userapi/utils/logger"
	"fmt"
)

var (
	userDB = map[string]*User{}
)

const (
	StatusActive           = "active"
	findByEmailandPassword = "SELECT id, first_name, last_name, email, status, password, date_created FROM users WHERE email=? AND password=? AND status=?;"
	queryInsertUser        = "INSERT INTO users(id, first_name, last_name, date_created, email, status, password) VALUES(?,?,?,?,?,?,?);"
	queryGetUserByID       = "SELECT id, first_name, last_name, email, status, password, date_created FROM users WHERE id=?;"
	findByStatus           = "SELECT id, first_name, last_name, email, status, password, date_created FROM users WHERE status=?;"
	queryUpdateUser        = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser        = "DELETE FROM users WHERE id=?"
)

func (input *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewRestErrInteralServer("Internal server error")
	}
	defer stmt.Close()
	result, err := stmt.Exec(input.Id, input.Firstname, input.Lastname, input.Datecreated, input.Email, input.Status, input.Password)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	input.Id, err = result.LastInsertId()
	return nil
}

func (input *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(input.Id)
	if getErr := result.Scan(&input.Id, &input.Firstname, &input.Lastname, &input.Email, &input.Status, &input.Password, &input.Datecreated); getErr != nil {
		return errors.NewRestErrInteralServer(getErr.Error())
	}
	return nil
}

func (input *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(input.Id)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	return nil
}

func (input *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(&input.Firstname, &input.Lastname, &input.Email, &input.Id)
	if err != nil {
		return errors.NewRestErrInteralServer(err.Error())
	}
	return nil
}

func (input *User) FindByStatus() ([]User, *errors.RestErr) {
	results := make([]User,0)
	stmt, err := users_db.Client.Prepare(findByStatus)
	if err != nil {
		return nil, errors.NewRestErrInteralServer(err.Error())
	}
	defer stmt.Close()
	users, err := stmt.Query(input.Status)
	if err != nil {
		log.Log.Info(err.Error())
		return nil, errors.NewRestErrInteralServer(err.Error())
	}
	for users.Next(){
		var user User
		err = users.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Status, &user.Password, &user.Datecreated)
		if err != nil {
			return nil, errors.NewRestErrInteralServer(err.Error())
		}
		results = append(results, user)
	}
	log.Log.Info(results)
	if len(results) == 0 {
		return nil, errors.NewRestErrResourceNotFound(fmt.Sprintf("No record found with status %s", input.Status))
	}
	return results, nil
}
