package api

import (
	"github.com/artshpakov/grapho/src/models"
	"github.com/gin-gonic/gin"
)

func UsersIndex(c *gin.Context) {
	users := make([]models.User, 0)
	c.JSON(200, &users)
}
