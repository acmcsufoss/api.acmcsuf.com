package main

import (
	// "context"
	// "database/sql"
	// "fmt"
	"log"
	// "net/http"
	// "os"
	// "os/signal"
	// "syscall"

	"github.com/gin-gonic/gin"

	// "github.com/acmcsufoss/api.acmcsuf.com/internal/api"
	// "github.com/acmcsufoss/api.acmcsuf.com/internal/db/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
