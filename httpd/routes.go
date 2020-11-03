package main

import (
	"lillyAppBackend/httpd/handlers"
	"net/http"
)

type handlerInfo struct {
	route       string
	handlerFunc func(http.ResponseWriter, *http.Request)
	corsMethods []string
}

func openRoutes() []handlerInfo {
	var routes []handlerInfo

	routes = []handlerInfo{
		{
			route:       "/images",
			handlerFunc: handlers.Index,
			corsMethods: []string{http.MethodGet, http.MethodOptions},
		},
		{
			route:       "/images/{imageName}",
			handlerFunc: handlers.Show,
			corsMethods: []string{http.MethodGet, http.MethodOptions},
		},
		{
			route:       "/images",
			handlerFunc: handlers.Create,
			corsMethods: []string{http.MethodPost, http.MethodOptions},
		},
	}

	return routes
}
