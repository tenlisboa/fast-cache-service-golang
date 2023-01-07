package main

import (
    "github.com/gin-gonic/gin"
    routes "github.com/tenlisboa/cache_service/application"
    "log"
)

func main() {
    r := gin.Default()

    routes.Routes(r)

    if err := r.Run(); err != nil {
        log.Fatal(err)
    }
}