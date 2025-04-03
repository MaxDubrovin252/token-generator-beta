package handler

import (
	"token-generator/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) Handler {
	return Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	token := router.Group("/token")
	{
		token.POST("/", h.PostHandler)
		token.POST("/get", h.GetMessage)

	}
	return router
}
