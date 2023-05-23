package controllers

import (
	"encoding/json"
	"errors"
	"gin_redis_rest/models"
	"gin_redis_rest/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func GetVersionCtl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.ApiRespondNoData{
			Success: true,
			Message: "Version 1.0.0",
		})
		return
	}
}

func GetKeyCtl() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("key")
		resp := services.GetKeyService(key)
		if resp.Success {
			c.JSON(http.StatusOK, resp)
			return
		} else {
			c.AbortWithError(400, errors.New(resp.Message))
			return
		}
	}
}

func DeleteKeyCtl() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Param("key")
		resp := services.DeleteKeyService(key)
		if resp.Success {
			c.JSON(http.StatusOK, resp)
			return
		} else {
			c.AbortWithError(400, errors.New(resp.Message))
			return
		}
	}
}

func PutKeyCtl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var keyData models.RedisDataModel

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		err = json.Unmarshal(body, &keyData)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		resp := services.PutKeyService(keyData)
		if resp.Success {
			c.JSON(http.StatusOK, resp)
			return
		} else {
			c.AbortWithError(400, errors.New(resp.Message))
			return
		}
	}
}
