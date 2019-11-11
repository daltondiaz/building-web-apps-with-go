package main

import (
	"log"
	"net/http"
	"path"
	"text/template"

	"gopkg.in/unrolled/render.v1"
)

func main() {

	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		// Assumes you have a template in ./templates called "example.tmpl"
		// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		r.HMTL(w, http.StatusOK, "example", nil)
	})

	// http.HandleFunc("/", ShowBooks)
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
