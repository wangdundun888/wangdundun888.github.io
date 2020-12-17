package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:w1215d1215c@/wdd?charset=utf8")
}

func main() {

	stmt, _ := db.Prepare("insert into user set username=?,password=?")
	stmt.Exec("wdd", "123456")

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var username, password string
		err := rows.Scan(&id, &username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, " ", username, " ", password)
	}
}
