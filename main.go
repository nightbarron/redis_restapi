package main

import (
	"gin_redis_rest/helpers"
	"gin_redis_rest/routers"
)

func main() {

	// init log
	helpers.InitLogger()

	r := routers.SetupRouter()
	r.Run(":8080")

}
