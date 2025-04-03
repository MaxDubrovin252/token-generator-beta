package handler

import (
	"token-generator/internal/generator"
	"token-generator/pkg/entity"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Request struct {
	Message string `json:"message"`
}

const tokenLencht = 16

func (h *Handler) PostHandler(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	token, err := generator.NewRandomString(tokenLencht)

	h.services.CreateToken(req.Message, token)
	if err != nil {

		logrus.Info("cannot generate new token", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"token": token})

}

func (h *Handler) GetMessage(c *gin.Context) {
	var req entity.Data

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	message, err := h.services.GetMessage(req.Token)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{
		"message": message,
	})

}
