package main

import (
	"context"
	"github.com/gorilla/mux"
	log2 "github.com/h3isenbug/url-common/pkg/services/log"
	"github.com/h3isenbug/url-query/config"
	"github.com/h3isenbug/url-query/handlers"
	"github.com/h3isenbug/url-query/handlers/url"
	urlService "github.com/h3isenbug/url-query/services/url"
	"net/http"
)

func provideHTTPServer(router *mux.Router) (*http.Server, func()) {
	server := &http.Server{
		Addr:    ":" + config.Config.Port,
		Handler: router,
	}
	return server, func() { server.Shutdown(context.Background()) }
}

func provideURLQueryHandler(urlService urlService.Service, logService log2.LogService) url.QueryHandler {
	return url.NewQueryHandlerV1(urlService, logService)
}

func provideMuxRouter(queryHandler url.QueryHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.GorillaMuxURLParamMiddleware)
	router.Methods("GET").Path("/r/{short_path}").HandlerFunc(queryHandler.GetLongURL)
	return router
}
