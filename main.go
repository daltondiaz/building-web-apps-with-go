package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

// HelloWorld Method
func HelloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(res, "Hello World")
}

// App complete handler
func App() http.Handler {
	n := negroni.Classic()

	m := func(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
		fmt.Fprint(response, "before...")
		next(response, request)
		fmt.Fprint(response, "...after")
	}

	n.Use(negroni.HandlerFunc(m))

	r := httprouter.New()

	r.GET("/", HelloWorld)
	n.UseHandler(r)

	return n
}

func main() {
	http.ListenAndServe(":3000", App())
}
