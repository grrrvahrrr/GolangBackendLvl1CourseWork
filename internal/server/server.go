package server

import (
	"CourseWork/internal/config"
	"CourseWork/internal/dbbackend"
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
	ds  *dbbackend.DataStorage
}

func NewServer(addr string, h http.Handler, config config.Config) *Server {
	s := &Server{}

	//Server settings should come from config
	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.WriteTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.ReadHeaderTimeout) * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err := s.srv.Shutdown(ctx)
	if err != nil {
		//Log it
		log.Println(err)
	}
	cancel()
}

func (s *Server) Start(ds *dbbackend.DataStorage) {
	s.ds = ds
	// TODO: migrations
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			//Log it
			log.Println(err)
		}
	}()
}
