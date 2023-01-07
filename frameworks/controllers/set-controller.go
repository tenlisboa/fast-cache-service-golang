package controllers

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "io"
)

type SetRequest struct {
    Key string `json:"key"`
    Data any `json:"data"`
}
func SetController(c *gin.Context) {
    jsonData, err := io.ReadAll(c.Request.Body)
    if (err != nil) {
        c.AbortWithStatusJSON(400, gin.H{
            "error": true,
            "message": "It was not possible to store the data",
            })
    }
    var request SetRequest
    if err := json.Unmarshal(jsonData, &request); err != nil {
        c.AbortWithStatusJSON(400, gin.H{
            "error": true,
            "message": "Data is invalid",
            })
    }

//    cache.Set(request.Key, request.Data)

    c.JSON(200, gin.H{
        "error": false,
        "message": "Data inserted successfully!",
    })
}