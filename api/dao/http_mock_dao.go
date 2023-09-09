package dao

import (
	"database/sql"
	"github.com/jspmarc/mockha/model"
)

type HttpMockDao interface {
	Save(mock *model.HttpMock) (*model.HttpMock, error)
	Update(mock *model.HttpMock) (*model.HttpMock, error)
	DeleteById(id int64) error
	FindByGroup(group sql.NullString) ([]*model.HttpMock, error)
	FindAll() ([]*model.HttpMock, error)
}
