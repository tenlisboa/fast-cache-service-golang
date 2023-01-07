package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/tenlisboa/cache_service/frameworks/controllers"
)

func Routes(router *gin.Engine) {
    router.GET("/get/:key", controllers.GetController)

    router.POST("/set", controllers.SetController)
}