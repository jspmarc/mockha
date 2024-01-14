package service

import (
	"database/sql"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/dto/http_mock"
	"github.com/jspmarc/mockha/model"
)

type HttpMockService interface {
	Start() error
	RegisterMock(createRequest *http_mock.CreateRequest) (*http_mock.Response, error)
	EditMock(mock *model.HttpMock) (*model.HttpMock, error)
	DeleteMock(group sql.NullString, path string, method constants.HttpMethod) error
	GetAllMocks() ([]*model.HttpMock, error)
	GetMocksByGroup(group sql.NullString) ([]*model.HttpMock, error)
	GetMock(group sql.NullString, path string, method constants.HttpMethod) (*model.HttpMock, error)
	Stop() error
}
