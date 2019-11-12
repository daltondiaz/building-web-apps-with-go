package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	expected := "Hello Worrld"
	current := res.Body.String()

	if expected != current {
		t.Fatalf("Expected %s get %s", expected, current)
	} else {
		fmt.Println("Success")
	}
}
