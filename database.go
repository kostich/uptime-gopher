package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		strconv.Itoa(params.Port) + ")/" + params.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		return fmt.Errorf("can't open db conn: %v", err)
	}

	return nil
}

// Checks if required tables exist in the database.
func checkTables(params *dbParams) error {
	dsn := params.User + ":" + params.Password + "@tcp(" + params.Host + ":" +
		strconv.Itoa(params.Port) + ")/" + params.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	stmt := "SHOW TABLE STATUS FROM `" + params.Name + "`;"
	rows, err := db.Raw(stmt).Rows()
	defer rows.Close()

	// will store all results from the statement above
	type tablesResult struct {
		name, engine string
		version      int
		rowFormat    string
		rows, avgRowLen, dataLen, maxDataLen,
		indexLen, dataFree, aInc int64
		createTime, updTime, checkTime, collation,
		checksum, creatOptions, comment string
	}

	// but we just need the names of the tables
	var tableNames []string
	for rows.Next() {
		// SHOW TABLE STATUS returns rows with 18 fields
		var tables tablesResult
		rows.Scan(&tables.name, &tables.engine, &tables.version, &tables.rowFormat,
			&tables.rows, &tables.avgRowLen, &tables.dataLen, &tables.maxDataLen,
			&tables.indexLen, &tables.dataFree, &tables.aInc, &tables.createTime,
			&tables.updTime, &tables.checkTime, &tables.collation, &tables.checksum,
			&tables.creatOptions, &tables.comment)
		tableNames = append(tableNames, tables.name)
	}

	requiredTables := []string{"keywords", "pings", "ports", "web_requests"}
	if !reflect.DeepEqual(tableNames, requiredTables) {
		return fmt.Errorf("required db tables do not exist")
	}

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
