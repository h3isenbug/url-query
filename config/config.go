package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type config struct {
	Port string `env:"PORT"`

	DSN string `env:"DATABASE_URL"`

	RedisServer   string `env:"REDIS_SERVER"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB"`

	RabbitQueueName string `env:"RABBIT_QUEUE_NAME"`
	RabbitServer    string `env:"RABBIT_SERVER"`
}

var Config config

func init() {
	t := reflect.TypeOf(Config)
	v := reflect.ValueOf(&Config).Elem()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("env")
		stringValue, found := os.LookupEnv(tag)
		if !found {
			panic(fmt.Sprintf("environment variable %s not set", tag))
		}

		switch t.Field(i).Type.Kind() {
		case reflect.String:
			v.Field(i).SetString(stringValue)
		case reflect.Slice:
			byteArrayValue, err := base64.StdEncoding.DecodeString(stringValue)
			if err != nil {
				panic(fmt.Sprintf("environment variable %s has incorrect value. expected base64.", tag))
			}
			v.Field(i).SetBytes(byteArrayValue)
		case reflect.Int:
			intValue, err := strconv.ParseInt(stringValue, 10, 32)
			if err != nil {
				panic(fmt.Sprintf("environment variable %s has incorrect value. expected int.", tag))
			}
			v.Field(i).SetInt(intValue)
		default:
			panic("unknown config field type")
		}
	}
}
