package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//returns a pointer to a db connection
func connectToDB() (db *sql.DB, err error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/ipaas?parseTime=true&charset=utf8mb4", conf.DbUsername, conf.DbPassword, conf.DbHost, conf.DbPort))
}
