package main

import "net/http"

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to
// a controller.
type Action func(rw http.ResponseWriter, r *http.Request) error

// AppController - This is our Base Controller
type AppController struct{}

// Action - this function helps with error handling in a controller
func (controller *AppController) Action(action Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := action(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}
