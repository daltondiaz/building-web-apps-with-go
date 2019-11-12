package main

import (
	"fmt"
	"net/http"
)

// HelloWorld Method
func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":3000", nil)
}
