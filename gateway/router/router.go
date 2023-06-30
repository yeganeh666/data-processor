package router

import (
	"IofIPOS/gateway/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Register(handler *handlers.ObjectHandlerImpl) {
	route := gin.Default()
	route.Use(func(c *gin.Context) {
		c.Next()
	})

	r := route.Group("/api")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/pre-upload", handler.HandlePreUpload)

	// Start the server
	if err := route.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
