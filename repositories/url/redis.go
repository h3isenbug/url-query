package url

import (
	"github.com/go-redis/redis"
)

type RedisCacheRepository struct {
	nextLayer   ReadRepository
	redisClient *redis.Client
}

func NewRedisCacheRepository(redisClient *redis.Client, nextLayer ReadRepository) ReadRepository {
	return &RedisCacheRepository{nextLayer: nextLayer, redisClient: redisClient}
}

func (repo RedisCacheRepository) GetLongURL(shortPath string) (string, error) {
	var longURL, err = repo.redisClient.Get(shortPath).Result()
	if err == nil {
		return longURL, nil
	}

	if err != redis.Nil {
		return "", err
	}

	// if key is not in cache
	longURL, err = repo.nextLayer.GetLongURL(shortPath)
	if err != nil {
		return "", err
	}

	if err := repo.redisClient.Set(shortPath, longURL, 0).Err(); err != nil {
		return "", err
	}

	return longURL, nil
}
