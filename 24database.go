package main

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	defaultDB *sql.DB
	lock      = sync.Mutex{}
)

func main() {
	lock.Lock()
	defer lock.Unlock()
	if defaultDB == nil {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/chatroom")
		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err != nil {
			panic(err)
		}
		defaultDB = db
	}
}
