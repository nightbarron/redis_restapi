package services

import (
	"gin_redis_rest/configs"
	"gin_redis_rest/models"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func connectToRedis() (*redis.Client, error) {
	redisConfig, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisConfig.RedisHost + ":" + strconv.Itoa(redisConfig.RedisPort),
	})
	return rdb, nil
}

func PutKeyService(keyData models.RedisDataModel) models.ApiRespond {
	resp := models.ApiRespond{Success: true, Message: "Put key to Redis server", Data: keyData}
	rdb, err := connectToRedis()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Data = nil
		return resp
	}
	defer rdb.Close()
	log.Info(time.Second * time.Duration(keyData.Expire))
	err = rdb.Set(rdb.Context(), keyData.Key, keyData.Value, time.Second*time.Duration(keyData.Expire)).Err()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Data = nil
	}
	return resp
}

func GetKeyService(key string) models.ApiRespond {
	resp := models.ApiRespond{Success: true, Message: "Get key from Redis server"}
	rdb, err := connectToRedis()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Data = nil
		return resp
	}
	defer rdb.Close()

	val, err := rdb.Get(rdb.Context(), key).Result()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp
	}
	ttl, err := rdb.TTL(rdb.Context(), key).Result()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp
	}
	resp.Data = models.RedisDataModel{Key: key, Value: val, Expire: int(ttl / time.Second)}

	return resp
}

func DeleteKeyService(key string) models.ApiRespondNoData {
	resp := models.ApiRespondNoData{Success: true, Message: "Delete key from Redis server"}
	rdb, err := connectToRedis()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()

		return resp
	}
	defer rdb.Close()
	err = rdb.Del(rdb.Context(), key).Err()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
	}

	return resp
}
