package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var line []string
var id int
var num int
var date string
var data string

func main() {
	file, err := os.Open("fina_server_tutorial.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	dbUser := os.Getenv("SQL_DB_USER")
	dbPS := os.Getenv("SQL_DB_PS")

	login := dbUser + ":" + dbPS + "@/fina_db"

	db, err := sql.Open("mysql", string(login))
	if err != nil {
		log.Fatal("db error.")
	} else {
		fmt.Println("access")
	}
	defer db.Close()
	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		tmpid, _ := strconv.Atoi(line[0])
		tmpnum, _ := strconv.Atoi(line[1])

		id = tmpid
		num = tmpnum
		date = line[2]
		data = line[3]

		_, err := db.Exec(`INSERT INTO test (id, num,date,data) VALUES (?,?,?,?) `, id, num, date, data)
		if err != nil {
			panic(err)
		}
	}
}
