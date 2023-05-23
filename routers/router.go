package routers

import (
	ctl "gin_redis_rest/controllers"
	"gin_redis_rest/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := models.Routes{
		Router: gin.Default(),
	}

	r.Router.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for version 1
	apiV1 := r.Router.Group("/v1")
	apiV1.GET("/version", ctl.GetVersionCtl())
	apiV1.GET("/keys/:key", ctl.GetKeyCtl())

	apiV1.PUT("/keys", ctl.PutKeyCtl())

	apiV1.DELETE("/keys/:key", ctl.DeleteKeyCtl())

	return r.Router

}
