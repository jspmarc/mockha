package service

import (
	"database/sql"
	"github.com/jspmarc/mockha/api/dao"
	"github.com/jspmarc/mockha/api/service"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/dto/http_mock"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils/mapper"
)

type HttpMockService struct {
	httpMockDao        dao.HttpMockDao
	requestResponseDao dao.HttpRequestResponseDao
}

func NewHttpMockService(mockDao dao.HttpMockDao, requestResponseDao dao.HttpRequestResponseDao) service.HttpMockService {
	svc := &HttpMockService{}

	svc.httpMockDao = mockDao
	svc.requestResponseDao = requestResponseDao

	return svc
}

func (s *HttpMockService) RegisterMock(createRequest *http_mock.CreateRequest) (*model.HttpMock, error) {
	var err error

	mock := mapper.CreateRequestToModelHttpMock(createRequest)
	if mock, err = s.httpMockDao.Save(mock); err != nil {
		return nil, err
	}

	rr := mapper.CreateRequestToModelHttpRequestResponse(createRequest, mock.Id)
	if _, err = s.requestResponseDao.Save(rr); err != nil {
		return nil, err
	}

	return mock, nil
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

func (s *HttpMockService) GetMock(group sql.NullString, path string, method constants.HttpMethod) (*model.HttpMock, error) {
	return nil, nil
}

func (s *HttpMockService) ExecuteMock(group sql.NullString, path string, method constants.HttpMethod) (interface{}, error) {
	return nil, nil
}
