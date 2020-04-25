package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	dbuserid       = "DBUSERID"
	dbuserpassword = "DBUSERPASSWORD"
	dbname         = "DBNAME"
	dbhost         = "DBHOST"
	dbport         = "DBPORT"
)

var (
	db_userid       = os.Getenv(dbuserid)
	db_userpassword = os.Getenv(dbuserpassword)
	db_name         = os.Getenv(dbname)
	db_host         = os.Getenv(dbhost)
	db_port         = os.Getenv(dbport)
	Client          *sql.DB
)

func init() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s", db_userid, db_userpassword, db_host, db_name)
	fmt.Println(DSN)
	var err error
	Client, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database configured succesfully")
}
