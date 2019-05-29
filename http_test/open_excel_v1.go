package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

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

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM test ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Database{}
	res := []Database{}
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
		res = append(res, emp)
	}
	//fmt.Println(res)
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Database{}
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
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)

	http.HandleFunc("/show", Show)
	http.ListenAndServe(":8080", nil)
}
