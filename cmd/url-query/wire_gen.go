// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"net/http"
)

// Injectors from inject_app.go:

func wireUp() (*http.Server, func(), error) {
	client, cleanup := provideRedisClient()
	db, cleanup2, err := provideSQLXConnection()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	readRepository, err := provideURLRepository(client, db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	channel, cleanup3, err := provideAMQPChannel()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	readRecipient, err := provideReadRecipient(channel)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service := provideURLService(readRepository, readRecipient)
	logService := provideLogService()
	queryHandler := provideURLQueryHandler(service, logService)
	router := provideMuxRouter(queryHandler)
	server, cleanup4 := provideHTTPServer(router)
	return server, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
