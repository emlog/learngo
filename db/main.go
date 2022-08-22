package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/test_note?parseTime=true&loc=Local")

	var myTime time.Time
	rows, err := db.Query("SELECT create_time from empyee limit 1")
	fmt.Println(time.Now())
	fmt.Printf("%s\n", time.Now())
	if rows.Next() {
		if err = rows.Scan(&myTime); err != nil {
			panic(err)
		}
	}

	fmt.Println(myTime)
}

type Note struct {
	title string
}
