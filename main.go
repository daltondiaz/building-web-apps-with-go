package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/unrolled/render.v1"
)

// MyController - My struct controller
type MyController struct {
	AppController
	*render.Render
}

// Index - method initial
func (controller *MyController) Index(rw http.ResponseWriter, request *http.Request) error {
	controller.JSON(rw, 200, map[string]string{"first_message": "Hello World"})
	return nil
}

func main() {
	db := NewDB()
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", ShowBooks(db))
}

// ShowBooks get books
func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var title, author string
		err := db.QueryRow("select title, author from books").Scan(&title, &author)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(rw, "The first book is '%s' by '%s'", title, author)
	})
}

// NewDB create new db and table
func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(title text, author text)")
	if err != nil {
		panic(err)
	}
	return db
}
