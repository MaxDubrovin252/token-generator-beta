package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SRV struct {
	server *http.Server
}

func (s *SRV) Start(port string, handler *gin.Engine) error {
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.server.ListenAndServe()
}
