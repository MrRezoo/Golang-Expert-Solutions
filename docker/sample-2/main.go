// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func dbHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user:password@tcp(db:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Fprintf(w, "Connected to MySQL version: %s", version)
}

func main() {
	http.HandleFunc("/", dbHandler)
	http.ListenAndServe(":8080", nil)
}
