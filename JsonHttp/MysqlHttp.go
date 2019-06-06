package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

//var tmpl = template.Must(template.ParseGlob("form/*"))

type Database struct {
	Id   int
	Num  int
	Date string
	Data string
}

//func Index(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	selDB, err := db.Query("SELECT * FROM test ORDER BY id DESC")
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Database{}
//	res := []Database{}
//	for selDB.Next() {
//		var id int
//		var num int
//		var date, data string
//		err = selDB.Scan(&id, &num, &date, &data)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Num = num
//		emp.Date = date
//		emp.Data = data
//		res = append(res, emp)
//	}
//fmt.Println(res)
//	tmpl.ExecuteTemplate(w, "Index", res)
//	defer db.Close()
//}

//func Show(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Database{}
//	for selDB.Next() {
//		var id int
//		var num int
//		var date, data string
//		err = selDB.Scan(&id, &num, &date, &data)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Num = num
//		emp.Date = date
//		emp.Data = data
//	}
//	tmpl.ExecuteTemplate(w, "Show", emp)
//	defer db.Close()
//}

//func readJson() Database {
//	b, err := ioutil.ReadFile("./test.json")
//	if err != nil {
//		fmt.Println(err)
//	}

//	var table []Database
//	json.Unmarshal(b, &table)
//	return table
//}

func main() {
	//log.Println("Server started on: http://localhost:8080")
	//http.HandleFunc("/", Index)

	//http.HandleFunc("/show", Show)
	//http.ListenAndServe(":8080", nil)

	b, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var table []Database
	json.Unmarshal(b, &table) // JSON 문서의 내용을 변환하여 table 저장
	fmt.Println(table)
}
