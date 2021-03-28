package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/core"
)

func handleRequests() {
	for _, elem := range core.Urls {
		http.HandleFunc(elem.Url, elem.Controller)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Сервер запущен")
	handleRequests()
}
