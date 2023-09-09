package dao

import (
	"database/sql"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
)

type HttpMockDao interface {
	Save(mock *model.HttpMock) (*model.HttpMock, error)
	Upsert(mock *model.HttpMock) (*model.HttpMock, error)
	DeleteById(id int64) error
	FindOne(group sql.NullString, path string, method constants.HttpMethod) (*model.HttpMock, error)
	FindOneById(id int64) (*model.HttpMock, error)
	FindByGroup(group sql.NullString) ([]*model.HttpMock, error)
	FindAll() ([]*model.HttpMock, error)
}
