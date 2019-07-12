package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Database struct {
	Id   int    `json:"Id"`
	Num  int    `json:"Num"`
	Date string `json:"Date"`
	Data string `json:"Data"`
}

const (
	PORT = ":8080"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func accessCode() string {
	dbUser := os.Getenv("SQL_DB_USER")
	//dbUser := "test"
	dbPS := os.Getenv("SQL_DB_PS")
	code := dbUser + ":" + dbPS + "@/fina_db"

	return code
}

func dbConn() (db *sql.DB, err error) {
	//	dbUser := os.Getenv("SQL_DB_USER")
	//	dbPS := os.Getenv("SQL_DB_PS")
	//login := dbUser + ":" + dbPS + "@/fina_db"
	//var err error

	login := accessCode()
	db, error := sql.Open("mysql", string(login))
	return db, error
}

func SelectSQL(searchID int) (Database, error) {
	//x := 6581
	var emp Database
	var id, num int
	var date, data string

	db, err := dbConn()

	if err != nil {
		log.Println("Wrong user, id")
		panic(err.Error())

		return emp, err
	}
	defer db.Close()

	selDB, err := db.Query("SELECT * FROM test WHERE id=?", searchID)
	if err != nil {
		log.Println("Wrong table")
		panic(err.Error())
		return emp, err
	}
	for selDB.Next() {

		err = selDB.Scan(&id, &num, &date, &data)
		if err != nil {
			return emp, err
		}
		return Database{Id: id, Num: num, Date: date, Data: data}, nil
	}
	return emp, nil
}

func ShowSQLNoTable(w http.ResponseWriter, r *http.Request) {

	var intID int

	vars := mux.Vars(r)
	stringID := vars["Id"]

	intID, _ = strconv.Atoi(stringID)

	emp, err := SelectSQL(intID)

	if err != nil {
		tmpl.ExecuteTemplate(w, "Error", err)
	} else if emp.Id == 0 {
		emp.Data = "データがありません。"
		tmpl.ExecuteTemplate(w, "NoData", emp)
	} else {
		tmpl.ExecuteTemplate(w, "Show", emp)
	}

}

func main() {

	log.Println("Server started on: http://localhost:8080")
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{Id}", ShowSQLNoTable)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
