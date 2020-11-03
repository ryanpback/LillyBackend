package main

import (
	h "lillyAppBackend/helpers"
	"lillyAppBackend/httpd/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func runApp() {
	r := mux.NewRouter()

	for _, handlerOpts := range openRoutes() {
		r.HandleFunc(handlerOpts.route, middleware.HandleCors(handlerOpts.handlerFunc)).Methods(handlerOpts.corsMethods...)
	}

	h.LogInfo("Starting web api at port " + appConfig.AppPort)

	h.LogFatal(http.ListenAndServe(":"+appConfig.AppPort, r))
}
