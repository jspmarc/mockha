package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/josep/mockha/utils"
	"log"
	"net/http"
)

type MockService struct {
	port uint16
	srv  *http.Server
}

func NewMockService(port *uint16) *MockService {
	svc := &MockService{}

	if port == nil {
		port = new(uint16)
		*port = utils.GetRandomTcpPort()
	}
	svc.port = *port

	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.callMock)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", svc.port),
		Handler: mux,
	}
	svc.srv = srv

	return svc
}

func (s *MockService) RegisterMock() {
	println("test")
}

func (s *MockService) DeleteMock() {
}

func (s *MockService) EditMock() {
}

func (s *MockService) GetMock() {
}

func (s *MockService) callMock(w http.ResponseWriter, req *http.Request) {
	log.Println("A mock is called with path", req.URL.Path, "and method", req.Method)
}

func (s *MockService) Start() {

	go func() {
		log.Printf("A mock HTTP server started at %d\n", s.port)
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("Error when starting mock HTTP server", err)
		}
	}()
}

func (s *MockService) Stop() {
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Println("Error when shutting down mock HTTP server", err)
	}
}
