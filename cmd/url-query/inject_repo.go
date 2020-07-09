package main

import (
	"github.com/go-redis/redis"
	"github.com/h3isenbug/url-query/config"
	"github.com/h3isenbug/url-query/repositories/url"
	"github.com/jmoiron/sqlx"
)

func provideURLRepository(redisClient *redis.Client, con *sqlx.DB) (url.ReadRepository, error) {
	pgRepo, err := url.NewPostgresReadRepositoryV1(con)
	if err != nil {
		return nil, err
	}
	return url.NewRedisCacheRepository(redisClient, pgRepo), nil
}

func provideSQLXConnection() (*sqlx.DB, func(), error) {
	var con, err = sqlx.Open("postgres", config.Config.DSN)
	return con, func() { con.Close() }, err
}

func provideRedisClient() (*redis.Client, func()) {
	var client = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisServer,
		Password: config.Config.RedisPassword,
		DB:       config.Config.RedisDB,
	})
	return client, func() {
		client.Close()
	}
}
