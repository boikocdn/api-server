package main

import (
	"github.com/gin-gonic/gin"
	"cdn-server/router"
)

func main() {
	r := gin.Default()
	router.ConfigRouter(r)
	r.Run()
}
