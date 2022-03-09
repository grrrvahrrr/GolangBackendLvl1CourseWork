package server

import (
	"CourseWork/internal/dbbackend"
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	srv http.Server
	ds  *dbbackend.DataStorage
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	//Server settings should come from config
	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err := s.srv.Shutdown(ctx)
	if err != nil {
		log.Error(err)
	}
	cancel()
}

func (s *Server) Start(ds *dbbackend.DataStorage) {
	s.ds = ds
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
