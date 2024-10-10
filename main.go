package main

import (
    "github.com/gin-gonic/gin"
    "Portofolio/handlers"
)

func main() {
    router := gin.Default()

    // Serve static files
    router.Static("/static", "./static")

    // Load HTML templates
    router.LoadHTMLGlob("templates/*")

    // Define routes
    router.GET("/", handlers.Home)

    // Start server
    router.Run(":8080")
}
