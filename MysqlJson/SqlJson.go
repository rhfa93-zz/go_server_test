package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Id   int
	Num  int
	Date string
	Data string
}

func dbConn() (db *sql.DB) {
	dbUser := os.Getenv("SQL_DB_USER")
	dbPS := os.Getenv("SQL_DB_PS")
	login := dbUser + ":" + dbPS + "@/fina_db"
	db, err := sql.Open("mysql", string(login))
	if err != nil {
		log.Fatal("db error.")
	} else {
		fmt.Println("access")
	}
	return db
}

func main() {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM test ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Database{}
	var table []Database
	for selDB.Next() {
		var id int
		var num int
		var date, data string
		err = selDB.Scan(&id, &num, &date, &data)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Num = num
		emp.Date = date
		emp.Data = data

		table = append(table, emp)

	}
	file, _ := json.MarshalIndent(table, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
	//fmt.Println(file.)
	defer db.Close()
}
