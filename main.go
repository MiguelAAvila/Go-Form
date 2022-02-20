package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Thrid party package
)

func setUpDB() (*sql.DB, error) {
	//provide credentials for database
	const (
		host     = "localhost"
		port     = 5432
		user     = "library"
		password = "admin"
		dbname   = "library"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Establish a connection to the database
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}
	//test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}

//Dependencies (Things/variables)
//Dependency Injection (passing dependencies)
type application struct {
	db *sql.DB
}

func main() {
	var db, err = setUpDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	app := &application{
		db: db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/book/create", app.createBookForm)
	mux.HandleFunc("/book-add", app.createBook)
	log.Println("Starting server on :4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
