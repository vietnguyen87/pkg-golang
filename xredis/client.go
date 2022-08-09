package xredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"strings"
)

type Client struct {
	*redis.Client
}

type Config struct {
	SingleMode bool
	URL        string
	Password   string
}

func NewClient(client *redis.Client) Cmdable {
	return &Client{
		Client: client,
	}
}

func GetRedisClient(redisURL, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: password,
	})

	_, err := client.Ping(context.Background()).Result()
	return client, err
}

func GetRedisClusterClient(redisURLs string, password string) (*redis.ClusterClient, error) {
	addrs := strings.Split(redisURLs, ",")
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})

	_, err := client.Ping(context.Background()).Result()
	return client, err
}

type Closable interface {
	Close() error
}

func CloseRedis(c Closable) func() {
	return func() {
		err := c.Close()
		if err != nil {
			logrus.Errorf("Error: failed to close redis, error: %v", err)
		}
	}
}

func Init(config Config) (Cmdable, func(), error) {
	if config.SingleMode {
		redisClient, err := GetRedisClient(config.URL, config.Password)
		client := NewClient(redisClient)
		return client, CloseRedis(client), err
	}
	redisClient, err := GetRedisClusterClient(config.URL, config.Password)
	client := NewClusterClient(redisClient)
	return client, CloseRedis(client), err
}
