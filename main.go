package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/", HomeHandler)

	router.GET("/posts", PostsIndexHandler)
	router.POST("/posts", PostsCreateHandler)

	router.GET("/posts/:id", PostShowHandler)
	router.PUT("/posts/:id", PostUpdateHandler)
	router.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on: " + port)
	router.NotFound = http.FileServer(http.Dir("public"))
	http.ListenAndServe(":"+port, router)
}

// HomeHandler Redirect to home
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Home")
}

// PostsIndexHandler posts index
func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts index")
}

// PostsCreateHandler create new post
func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts create")
}

// PostShowHandler show post by id
func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

// PostUpdateHandler Update post
func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

// PostDeleteHandler Delete post
func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

// PostEditHandler Edit post
func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}
