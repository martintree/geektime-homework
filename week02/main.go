package main

import (
	"week02/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", service.TestUser)
	http.ListenAndServe(":8080", nil)
}
