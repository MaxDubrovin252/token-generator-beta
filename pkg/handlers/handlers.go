package handlers

import (
	"log/slog"
	"token-generator/logger/sl"
	"token-generator/pkg/generator"
	"token-generator/pkg/resp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TokenGenerate interface {
	NewToken(message string, token string) (int64, error)
	GetMessage(token string) (string, error)
}
type Request struct {
	Message string `json:"message" validate:"message"`
	Token   string `json:"token"`
}

type Response struct {
	resp.Response
	Token string `json:"token"`
}

const tokenLenght = 35

func PostHandler(log *slog.Logger, stor TokenGenerate) gin.HandlerFunc {
	return func(c *gin.Context) {
		const op = "handlers.Post"

		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("cannot decode request", sl.Err(err))
			c.JSON(400, resp.Error("cannot read your message"))
		}

		if err := validator.New().Struct(&req); err != nil {

			log.Error("invalid input", sl.Err(err))
			c.JSON(400, resp.Error("invalid input"))

		}
		token := req.Token

		if token == "" {
			token = generator.NewRandomString(tokenLenght)
		}
		log.Info("request decoded", slog.Any("request", req))
		responseOK(c, token)
	}

}

func responseOK(c *gin.Context, token string) {
	c.JSON(200, Response{
		Response: *resp.OK(),
		Token:    token,
	})
}
