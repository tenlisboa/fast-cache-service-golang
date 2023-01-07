package controllers

import "github.com/gin-gonic/gin"

func GetController(c *gin.Context) {
    key := c.Param("key")

//    data, _ := cache.Get(key)

    c.JSON(200, key)
}