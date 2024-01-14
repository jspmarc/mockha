package repository

import (
	"github.com/jspmarc/mockha/model"
)

type HttpMockRepository interface {
	Save(mock *model.HttpMock) (*model.HttpMock, error)
	Update(mock *model.HttpMock) (*model.HttpMock, error)
	DeleteById(id int64) error
	FindByGroup(group string) ([]*model.HttpMock, error)
	FindAll() ([]*model.HttpMock, error)
}
