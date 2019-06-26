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

func dbConn() (db *sql.DB) {
	dbUser := os.Getenv("SQL_DB_USER")
	dbPS := os.Getenv("SQL_DB_PS")
	login := dbUser + ":" + dbPS + "@/fina_db"
	db, err := sql.Open("mysql", string(login))
	if err != nil {
		log.Fatal("db error.")
	}
	return db
}

func SelectSQL(searchID int) (Database, error) {
	//x := 6581
	var emp Database
	var id, num int
	var date, data string

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM test WHERE id=?", searchID)
	if err != nil {
		return emp, err
	}
	defer db.Close()
	for selDB.Next() {

		err = selDB.Scan(&id, &num, &date, &data)
		if err != nil {
			return emp, err
		}
		return Database{Id: id, Num: num, Date: date, Data: data}, nil
		//emp.Id = id
		//emp.Num = num
		//emp.Date = date
		//emp.Data = data

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
