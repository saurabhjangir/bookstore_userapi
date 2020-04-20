package users_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

const (
	db_userid = "saurabh"
	db_userpassword = "January.2020"
	db_name = "usersDB"
	db_host = "127.0.0.1"
	db_port = "3306"
)

var Client *sql.DB

func init(){
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s",db_userid, db_userpassword,db_host, db_name)
	var err error
	Client , err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database configured succesfully")
}

