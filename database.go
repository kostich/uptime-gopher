package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type dbParams struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
}

// Checks if it's possible to establish connection to the database with
// parameters defined in conf.json file.
func checkDbConn(params *dbParams) error {
	dsn := params.User + ":" + params.Password + "@tcp(" + params.Host + ":" +
		strconv.Itoa(params.Port) + ")/" + params.Name
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("can't open db conn: %v", err)
	}

	return nil
}

// Checks if required tables exist in the database.
func checkTables(params *dbParams) error {
	dsn := params.User + ":" + params.Password + "@tcp(" + params.Host + ":" +
		strconv.Itoa(params.Port) + ")/" + params.Name
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	//stmt := "SHOW TABLE STATUS FROM `" + params.Name + "`"
	//rows, err := db.Query(stmt)
	//if err != nil {
	//	return fmt.Errorf("can't get data from db: %v", err)
	//}
	// TODO: check rows and see if all the tables are there and that the
	// table structure is correct.

	return nil
}

// Creates required tables in the database.
func createTables() {
	return
}

// Logs the data about the ping check to the database.
func logPing() {
	return
}

// Logs the data about the keyword check to the database.
func logKeyword() {
	return
}

// Logs the data about the port check to the database.
func logPort() {
	return
}

// Logs the data about the request check to the database.
func logRequest() {
	return
}
