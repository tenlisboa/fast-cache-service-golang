package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/tenlisboa/cache_service/domains/usecases"
    "github.com/tenlisboa/cache_service/services"
)

func GetController(c *gin.Context) {
    key := c.Param("key")

    cacheService := services.GetCache()
    usecase := usecases.NewGetDataUsecase(cacheService)
    data, _ := usecase.Execute(usecases.GetDataInput{
        Key: key,
    })

    c.JSON(200, data)
}