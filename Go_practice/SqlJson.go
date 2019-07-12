//Mysql to Json file

package main

import (
	"database/sql"
	"encoding/json"
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

func accessCode() string {
	dbUser := os.Getenv("SQL_DB_USER")
	dbPS := os.Getenv("SQL_DB_PS")

	code := dbUser + ":" + dbPS + "@/fina_db"

	return code
}

func dbConn() (db *sql.DB, err error) {
	//dbUser := os.Getenv("SQL_DB_USER")
	//dbPS := os.Getenv("SQL_DB_PS")
	//login := dbUser + ":" + dbPS + "@/fina_db"
	login := accessCode()

	db, error := sql.Open("mysql", string(login))
	//if err != nil {
	//	log.Fatal("db error.")
	//} else {
	//	log.Println("access")
	//}
	return db, error
}

func main() {
	db, err := dbConn()

	if err != nil {
		log.Println("Wrong user, id")
		panic(err.Error())

		return emp, err
	}
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM test ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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

}
