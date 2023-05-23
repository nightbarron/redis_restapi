package models

import "github.com/gin-gonic/gin"

type Config struct {
	RedisHost string `json:"redis_host"`
	RedisPort int    `json:"redis_port"`
}

type Routes struct {
	Router *gin.Engine
}

type ApiRespond struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApiRespondNoData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RedisDataModel struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Expire int    `json:"expire" default:"0"`
}
