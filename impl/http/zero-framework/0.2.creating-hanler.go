package main

import (
	"fmt"
	"net/http"
)

// 002 - Creating a handler

// defining a port
const port = 8080

// handler structure/ could be anything, but needs to implement http.Handler interface
type myHandler struct{}


// all handlers must have this method
func (m *myHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("did it"))
}

func main(){
	http.Handle("/", &myHandler{})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}