package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jspmarc/mockha/api/repository"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/dto/http_mock"
	"github.com/jspmarc/mockha/entities"
	"github.com/jspmarc/mockha/model"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

type httpMockServiceServer struct {
	srv *http.Server
	wg  *sync.WaitGroup
}

type HttpMockService struct {
	httpMockDao        repository.HttpMockRepository
	requestResponseDao repository.HttpRequestResponseRepository
	server             httpMockServiceServer
}

func NewHttpMockService(mockDao repository.HttpMockRepository, requestResponseDao repository.HttpRequestResponseRepository, mockServerAddress string) service.HttpMockService {
	svc := &HttpMockService{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", executeMock)

	srv := &http.Server{Addr: mockServerAddress, Handler: mux}
	wg := &sync.WaitGroup{}

	svc.httpMockDao = mockDao
	svc.requestResponseDao = requestResponseDao
	svc.server = httpMockServiceServer{srv, wg}

	return svc
}

func (s *HttpMockService) Start() error {
	server := s.server

	srv := server.srv
	wg := server.wg

	go func() {
		defer wg.Done()

		log.Info().Msgf("Starting mock HTTP server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().
				Err(err).
				Msg("Unable to start HTTP mock server")
		}
	}()

	wg.Add(1)

	return nil
}

func (s *HttpMockService) RegisterMock(headers *entities.Headers, createRequest *http_mock.CreateRequest) (*http_mock.Response, error) {
	var err error

	mock := createRequest.ToModelHttpMock()
	if mock, err = s.httpMockDao.Save(mock); err != nil {
		log.Error().
			Str("requestId", headers.RequestId).
			Err(err).
			Msg("Got error when saving mock to DB")
		return nil, err
	}

	rr := createRequest.ToCreateRequestToModelHttpRequestResponse(mock.Id)
	if _, err = s.requestResponseDao.Save(rr); err != nil {
		return nil, err
	}

	log.Info().
		Str("requestId", headers.RequestId).
		Int64("mockId", mock.Id).
		Msg("Successfully created a new mock")

	return http_mock.NewHttpMockResponse(mock, rr), nil
}

func (s *HttpMockService) EditMock(mock *model.HttpMock) (*model.HttpMock, error) {
	return nil, nil
}

func (s *HttpMockService) DeleteMock(group sql.NullString, path string, method constants.HttpMethod) error {
	return nil
}

func (s *HttpMockService) GetAllMocks() ([]*model.HttpMock, error) {
	return nil, nil
}

func (s *HttpMockService) GetMocksByGroup(group sql.NullString) ([]*model.HttpMock, error) {
	return nil, nil
}

func (s *HttpMockService) GetMock(group sql.NullString, path string, method constants.HttpMethod) (*http_mock.Response, error) {
	return nil, nil
}

func (s *HttpMockService) Stop() error {
	server := s.server

	if err := server.srv.Shutdown(context.Background()); err != nil {
		return err
	}
	log.Info().Msg("Mock HTTP server stopped")
	server.wg.Wait()

	return nil
}

func executeMock(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Msg("Got mock HTTP request")
	w.Write([]byte("hello, world"))
}
