package service

import (
	"database/sql"
	"github.com/jspmarc/mockha/api/dao"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
	"log"
	"net/http"
)

type MockService struct {
	httpMock dao.HttpMockDao
}

func NewMockService(mockDao dao.HttpMockDao) service.MockService {
	svc := &MockService{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.callMock)

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

func (s *MockService) ExecuteMock(group sql.NullString, path string, method constants.HttpMethod) (interface{}, error) {
	return nil, nil
}

func (s *MockService) callMock(w http.ResponseWriter, req *http.Request) {
	log.Println("A mock is called with path", req.URL.Path, "and method", req.Method)
}
