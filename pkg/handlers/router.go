package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func InitRouter(log *slog.Logger, storage TokenGenerate) *gin.Engine {
	router := gin.New()

	router.POST("/token", PostHandler(log, storage))

	return router
}
