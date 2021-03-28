package core

import (
	"net/http"
	"restapi/controllers"
)

type pair struct {
	Url        string
	Controller func(http.ResponseWriter, *http.Request)
}

var Urls = []pair{
	pair{
		Url:        "/",
		Controller: controllers.HomePage,
	},
	pair{
		Url:        "/all_articles/",
		Controller: controllers.AllArticles,
	},
}
