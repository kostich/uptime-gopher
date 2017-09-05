package main

import (
	"fmt"
	"strconv"
	"time"

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

type keywordsTable struct {
	ID       uint
	Datetime time.Time
	Host     string
	Keyword  string
	State    bool
	Comment  string
}

type pingsTable struct {
	ID       uint
	Datetime time.Time
	Host     string
	State    bool
	Comment  string
}

type portsTable struct {
	ID       uint
	Datetime time.Time
	Host     string
	Port     int
	state    bool
	Comment  string
}

type webRequestsTable struct {
	ID          uint
	Datetime    time.Time
	Host        string
	DesiredResp int
	ActualResp  int
	Comment     string
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

	if !db.HasTable(&keywordsTable{}) {
		return fmt.Errorf("required db tables do not exist")
	} else if !db.HasTable(&pingsTable{}) {
		return fmt.Errorf("required db tables do not exist")
	} else if !db.HasTable(&portsTable{}) {
		return fmt.Errorf("required db tables do not exist")
	} else if !db.HasTable(&webRequestsTable{}) {
		return fmt.Errorf("required db tables do not exist")
	}

	return nil
}

// Creates required tables in the database.
func createTables(params *dbParams) error {
	dsn := params.User + ":" + params.Password + "@tcp(" + params.Host + ":" +
		strconv.Itoa(params.Port) + ")/" + params.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	db.CreateTable(&keywordsTable{})
	db.CreateTable(&pingsTable{})
	db.CreateTable(&portsTable{})
	db.CreateTable(&webRequestsTable{})

	return nil
}

// Logs the data about the ping check to the database.
func logPing(params *dbParams, data *pingResp) error {
	dsn := params.User + ":" + params.Password + "@tcp(" + params.Host + ":" +
		strconv.Itoa(params.Port) + ")/" + params.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	nr := pingsTable{Datetime: data.datetime, Host: data.host,
		State: data.state, Comment: data.comment}
	db.Create(&nr)

	return nil
}

// Logs the data about the keyword check to the database.
func logKeyword(params *dbParams, data *keywordResp) error {
	dsn := params.User + ":" + params.Password + "@tcp(" + params.Host + ":" +
		strconv.Itoa(params.Port) + ")/" + params.Name + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("can't get db handle: %v", err)
	}
	defer db.Close()

	nr := keywordsTable{Datetime: data.datetime, Host: data.host,
		Keyword: data.keyword, State: data.state, Comment: data.comment}
	db.Create(&nr)

	return nil
}

// Logs the data about the port check to the database.
func logPort() {
	return
}

// Logs the data about the request check to the database.
func logRequest() {
	return
}
