package main

import (
	"fmt"
	"net/http"
)

// Using Handle type
type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

//Using Handler
func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Something"))
}

func main() {
	fmt.Println("test")
	http.Handle("/foo", &fooHandler{Message: "foo called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
