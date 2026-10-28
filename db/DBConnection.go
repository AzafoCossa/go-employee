package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection()(database *sql.DB){
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "Tr@ficpr0gr@mmer9669!!11"
	databaseName := "go_employee"

	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)

	if error != nil{
		panic(error.Error())
	}

	return database
}