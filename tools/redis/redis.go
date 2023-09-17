package redis

import (
	"github.com/go-redis/redis/v8"
)

var (
	UserRedisClient   *redis.Client
	GroupRedisClient  *redis.Client
	RecallRedisClient *redis.Client
)

var (
	UserProfileKey = "_user_profile"
	GroupInfoKey   = "_group"
	RecallGroupKey = "_recall_group"
)

func CreateRedisClient() {
	createGroupRedisClient()
	createUserRedisClient()
	createRecallRedisClient()
}

func createRecallRedisClient() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果没有设置密码则为空字符串
		DB:       2,                // Redis 数据库索引
	})
	// Ping 服务器，检查连接是否成功
	_, err := client.Ping(client.Context()).Result()
	RecallRedisClient = client
	if err != nil {
		panic(err)
	}
}
func createGroupRedisClient() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果没有设置密码则为空字符串
		DB:       1,                // Redis 数据库索引
	})
	// Ping 服务器，检查连接是否成功
	_, err := client.Ping(client.Context()).Result()
	GroupRedisClient = client
	if err != nil {
		panic(err)
	}
}

func createUserRedisClient() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 服务器密码，如果没有设置密码则为空字符串
		DB:       0,                // Redis 数据库索引
	})
	// Ping 服务器，检查连接是否成功
	_, err := client.Ping(client.Context()).Result()
	UserRedisClient = client
	if err != nil {
		panic(err)
	}
}
