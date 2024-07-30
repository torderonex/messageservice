package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/torderonex/messageservice/internal/service"

	docs "github.com/torderonex/messageservice/docs"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")

	message := v1.Group("/messages")

	{
		message.POST("send", h.sendMessage)
		message.PUT("process", h.processMessages)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
