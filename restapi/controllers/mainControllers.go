package controllers

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Тестовый REST API")
	fmt.Println("Endpoint Hit: HomePage")
}
