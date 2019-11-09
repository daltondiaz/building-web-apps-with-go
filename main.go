package main

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

func main() {

	// Middleware stack
	/* n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyFirstGoMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Run(":8080")*/

	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":8080", nil)
}

// Book struct
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	// json:"-" not show the field
	Password string `json:"-"`
	// json:"-," show - instead Email
	Email string `json:"-,"`
	// omitempty , if empty not show lastname
	LastName string `json:"lastname,omitempty"`
}

// ShowBooks in json
func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{
		"12 Rules of Life: An Antitode to Caos",
		"Jordan Peterson",
		"123",
		"dalton@example.com",
		""}

	pathFiles := path.Join("templates", "index.html")
	goTemplate, err := template.ParseFiles(pathFiles)

	// js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := goTemplate.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// MyFirstGoMiddleware this is my first Middleware make in Go Lang
func MyFirstGoMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")

	if r.URL.Query().Get("password") == "123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back...")
}
