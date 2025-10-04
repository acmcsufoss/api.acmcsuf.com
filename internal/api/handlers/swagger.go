package handlers

import (
	docs "github.com/acmcsufoss/api.acmcsuf.com/docs"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggerHandler() gin.HandlerFunc {
	docs.SwaggerInfo.Title = "ACM CSUF API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "v1/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Description = utils.SwaggerDescription
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
