package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/models"
)

func AllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(models.Articles)
}
