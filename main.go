package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func getV1PingHandler(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "v1 pong",
        },
    )
}

func getV2PingHandler(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "v2 pong",
        },
    )
}

func main() {
    engine := gin.Default()

    v1RouterGroup := engine.Group("/v1")
    {
        v1RouterGroup.GET("/ping", getV1PingHandler)
    }

    v2RouterGroup := engine.Group("/v2")
    {
        v2RouterGroup.GET("/ping", getV2PingHandler)
    }

    engine.Run()
}
