package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/docs"
)

func NewSwaggerHandler() gin.HandlerFunc {
	docs.SwaggerInfo.Title = "ACM CSUF API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
