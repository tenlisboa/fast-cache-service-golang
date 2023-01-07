package helpers

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io"
)

func ParseBody[Type interface{}](context *gin.Context) Type {
    jsonData, err := io.ReadAll(context.Request.Body)
    if (err != nil) {
        context.AbortWithStatusJSON(400, gin.H{
            "error": true,
            "message": "It was not possible to store the data",
            })
    }
    var request Type
    if err := json.Unmarshal(jsonData, &request); err != nil {
        context.AbortWithStatusJSON(400, gin.H{
            "error": true,
            "message": "Data is invalid",
            })
    }

    return request
}