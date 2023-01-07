package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/tenlisboa/cache_service/domains/usecases"
    "github.com/tenlisboa/cache_service/frameworks/helpers"
    "github.com/tenlisboa/cache_service/services"
)

type SetRequest struct {
    Key string `json:"key"`
    Data any `json:"data"`
}
func SetController(c *gin.Context) {
    request := helpers.ParseBody[SetRequest](c)

    cacheService := services.GetCache()
    usecase := usecases.NewSetDataUsecase(cacheService)
    usecase.Execute(usecases.SetDataInput{
        Key: request.Key,
        Data: request.Data,
    })

    c.JSON(200, gin.H{
        "error": false,
        "message": "Data inserted successfully!",
    })
}