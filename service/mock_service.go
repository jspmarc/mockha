package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jspmarc/mockha/api/dao"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils"
	"log"
	"net/http"
)

type MockService struct {
	port     uint16
	srv      *http.Server
	httpMock dao.HttpMockDao
}

func NewMockService(mockDao dao.HttpMockDao, port *uint16) service.MockService {
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

	svc.httpMock = mockDao

	return svc
}

func (s *MockService) RegisterMock(mock *model.HttpMock) (*model.HttpMock, error) {
	println("test")

	return nil, nil
}

func (s *MockService) EditMock(mock *model.HttpMock) (*model.HttpMock, error) {
	return nil, nil
}

func (s *MockService) DeleteMock(group sql.NullString, path string, method constants.HttpMethod) error {
	return nil
}

func (s *MockService) GetAllMocks() ([]*model.HttpMock, error) {
	return nil, nil
}

func (s *MockService) GetMocksByGroup(group sql.NullString) ([]*model.HttpMock, error) {
	return nil, nil
}

func (s *MockService) GetMock(group sql.NullString, path string, method constants.HttpMethod) (*model.HttpMock, error) {
	return nil, nil
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

func (s *MockService) Stop() error {
	return s.srv.Shutdown(context.Background())
}
