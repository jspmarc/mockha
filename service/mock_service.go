package service

import (
	"database/sql"
	"github.com/jspmarc/mockha/api/dao"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
)

type MockService struct {
	httpMockDao        dao.HttpMockDao
	requestResponseDao dao.HttpRequestResponseDao
}

func NewMockService(mockDao dao.HttpMockDao, requestResponseDao dao.HttpRequestResponseDao) service.MockService {
	svc := &MockService{}

	svc.httpMockDao = mockDao
	svc.requestResponseDao = requestResponseDao

	return svc
}

func (s *MockService) RegisterMock(mock *model.HttpMock) (*model.HttpMock, error) {
	println("tests")

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
