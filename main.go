package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // JSON dönen endpoint
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    // Statik dosya servisi (örneğin ./static klasörü)
    r.Static("/static", "./static")

    // Server 8080 portunda başlar
    r.Run(":8080")
}
