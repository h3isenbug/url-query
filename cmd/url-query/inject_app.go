//+build wireinject

package main

import (
	"github.com/google/wire"
	"net/http"
)

func wireUp() (*http.Server, func(), error) {
	wire.Build(provideURLRepository,provideReadRecipient,provideURLService, provideURLQueryHandler, provideHTTPServer, provideAMQPChannel, provideLogService, provideMuxRouter, provideSQLXConnection, provideRedisClient)

	return &http.Server{}, func() {

	}, nil
}
