package service

import (
	"database/sql"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
)

type MockService interface {
	RegisterMock(mock *model.HttpMock) (*model.HttpMock, error)
	EditMock(mock *model.HttpMock) (*model.HttpMock, error)
	DeleteMock(group sql.NullString, path string, method constants.HttpMethod) error
	GetAllMocks() ([]*model.HttpMock, error)
	GetMocksByGroup(group sql.NullString) ([]*model.HttpMock, error)
	GetMock(group sql.NullString, path string, method constants.HttpMethod) (*model.HttpMock, error)
	Start()
	Stop() error
}
