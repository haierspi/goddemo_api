package redis

import (
	"context"

	"starfission_go_api/pkg/setting"

	"github.com/go-redis/redis/v8"
)

// 定义RabbitMQ对象
type Redis struct {
	Client *redis.Client
}

func NewConn(c *setting.RedisS) (*Redis, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.Db,
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}
	return &Redis{Client: client}, nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
