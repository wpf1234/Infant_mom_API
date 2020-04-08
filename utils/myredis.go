package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

func NewClient(host, pwd string, port, db, pool int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	RedisClient = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               addr,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           pwd,
		DB:                 db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           pool,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	if _, err := RedisClient.Ping().Result(); err != nil {
		log.Error("Connect redis failed: ", err)
		return err
	}
	return nil
}
