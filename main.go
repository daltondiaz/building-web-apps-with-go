package main

import (
	"net/http"

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
	controller := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":8080", controller.Action(controller.Index))
}
