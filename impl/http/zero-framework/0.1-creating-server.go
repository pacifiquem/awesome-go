package main

import (
	"fmt"
	"net/http"
)

const port = 8080

func main(){
	// jusr pull in http packeges and you're good to go
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}