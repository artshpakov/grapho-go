package main

import (
	"github.com/artshpakov/grapho/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", api.UsersIndex)

	r.Run(":3000")
}
