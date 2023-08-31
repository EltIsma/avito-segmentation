package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if err :=   s.httpServer.ListenAndServe(); err!= nil{
		return fmt.Errorf("error starting server: %w", err)
	}
	 return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err :=  s.httpServer.Shutdown(ctx); err!= nil{
		return fmt.Errorf("error starting server: %w", err)
	}
	 return nil
}