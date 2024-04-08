package main

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

type User struct {
	Name  string `db:"username"`
	Email string `db:"email"`
 }

func main() {
  //connect to a PostgreSQL database
  // Replace the connection details (user, dbname, password, host) with your own
    db, err := sqlx.Connect(dbDriver, dbSource)
    if err != nil {
        log.Fatalln(err)
    }
  
    defer db.Close()

    // Test the connection to the database
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }

}