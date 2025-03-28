package server

import (
	"net/http"
	"time"
)

type SRV struct {
	httpServer *http.Server
}

func (s *SRV) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    1 * time.Minute,
		MaxHeaderBytes: 1 << 15,
	}

	return s.httpServer.ListenAndServe()
}
