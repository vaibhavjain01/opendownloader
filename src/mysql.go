package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlUtility struct {
	Username string
	Password string
	DbName	string
	LogObj Logger
}

func (mySqlUtility MySqlUtility) Init () (*sql.DB) {
	db, err := sql.Open("mysql", mySqlUtility.Username + ":" + mySqlUtility.Password + "@tcp(127.0.0.1:3306)/" + mySqlUtility.DbName)
	if err != nil {
		panic(err.Error())
		return nil
    } else {
		mySqlUtility.LogObj.LogToConsole("Connection Established")
	}
	return db
}

func (mySqlUtility MySqlUtility) InsertIntoLinks(link string) {
	//db := mySqlUtility.Init()
	//insert, err := db.Query("INSERT INTO links VALUES ( " + link + ", 0, NOW())")

	db, err := sql.Open("mysql", mySqlUtility.Username + ":" + mySqlUtility.Password + "@tcp(127.0.0.1:3306)/" + mySqlUtility.DbName)
	if err != nil {
		panic(err.Error())
    } else {
		mySqlUtility.LogObj.LogToConsole("Connection Established")
	}
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(1000)
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM links")
	if err != nil {
        panic(err.Error())
	}
	
	res, err := stmt.Exec()
    if err != nil {
        panic(err.Error())
    } else {
		mySqlUtility.LogObj.LogToConsole("Row Inserted")
		//mySqlUtility.LogObj.LogToConsole(res)
		fmt.Println(res)
	}
	defer stmt.Close()
}