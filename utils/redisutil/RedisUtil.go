package redisutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jumpingcoder/quickutil4go/utils/logutil"
	_ "github.com/lib/pq"
	"time"
)

var redisClients map[string]*redis.Client = make(map[string]*redis.Client)

func InitFromConfig(configs []interface{}, decryptKey string, decryptHandler func(content string, decryptKey string) string) bool {
	for _, config := range configs {
		configMap := config.(map[string]interface{})
		if configMap["RedisName"] == nil || configMap["Addr"] == nil {
			logutil.Error("RedisName、Addr为必填项", nil)
			return false
		}
		redisName := configMap["RedisName"].(string)
		addr := configMap["Addr"].(string)
		password := ""
		if configMap["Password"] != nil {
			password = decryptHandler(configMap["Password"].(string), decryptKey)
		}
		db := 0
		if configMap["DB"] != nil {
			db = int(configMap["DB"].(float64))
		}
		maxRetries := -1
		poolSize := 200
		minIdleConns := 100
		dialTimeout := 10 * time.Second
		readTimeout := 30 * time.Second
		writeTimeout := 30 * time.Second
		poolTimeout := 10 * time.Second
		if configMap["MaxRetries"] != nil {
			maxRetries = int(configMap["MaxRetries"].(float64))
		}
		if configMap["PoolSize"] != nil {
			poolSize = int(configMap["PoolSize"].(float64))
		}
		if configMap["MinIdleConns"] != nil {
			minIdleConns = int(configMap["MinIdleConns"].(float64))
		}
		if configMap["DialTimeout"] != nil {
			dialTimeout = time.Duration(time.Second.Nanoseconds() * int64(configMap["DialTimeout"].(float64)))
		}
		if configMap["ReadTimeout"] != nil {
			readTimeout = time.Duration(time.Second.Nanoseconds() * int64(configMap["ReadTimeout"].(float64)))
		}
		if configMap["WriteTimeout"] != nil {
			writeTimeout = time.Duration(time.Second.Nanoseconds() * int64(configMap["WriteTimeout"].(float64)))
		}
		if configMap["PoolTimeout"] != nil {
			poolTimeout = time.Duration(time.Second.Nanoseconds() * int64(configMap["PoolTimeout"].(float64)))
		}
		options := &redis.Options{
			Addr:         addr,
			Password:     password,
			DB:           db,
			MaxRetries:   maxRetries,
			PoolSize:     poolSize,
			MinIdleConns: minIdleConns,
			DialTimeout:  dialTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			PoolTimeout:  poolTimeout,
		}
		AddRedisClient(redisName, options)
	}
	return true
}

var ctx = context.Background()

func AddRedisClient(redisName string, options *redis.Options) {
	if redisClients[redisName] != nil {
		logutil.Warn("Redis "+redisName+" already exists", nil)
	}
	client := redis.NewClient(options)
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		logutil.Error("Redis "+redisName+" ping failed", err)
	} else {
		logutil.Info("Redis "+redisName+" connected with "+pong, err)
	}
	redisClients[redisName] = client
}

func GetRedisClient(redisname string) *redis.Client {
	return redisClients[redisname]
}

func CloseRedisClient(redisName string) bool {
	if redisClients[redisName] == nil {
		logutil.Warn("Redis "+redisName+" not exists", nil)
		return true
	}
	err := redisClients[redisName].Close()
	redisClients[redisName] = nil
	if err != nil {
		logutil.Error("Redis "+redisName+" close failed", err)
	}
	return true
}
