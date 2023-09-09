package dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/api/dao"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
)

func NewHttpMockDao(db *sqlx.DB) dao.HttpMockDao {
	mockDao := &HttpMockDao{}

	mockDao.db = db

	return mockDao
}

type HttpMockDao struct {
	db *sqlx.DB
}

func (d *HttpMockDao) Save(mock *model.HttpMock) (*model.HttpMock, error) {
	return nil, nil
}

func (d *HttpMockDao) Upsert(mock *model.HttpMock) (*model.HttpMock, error) {
	return nil, nil
}

func (d *HttpMockDao) FindOne(group sql.NullString, path string, method constants.HttpMethod) (*model.HttpMock, error) {
	return nil, nil
}

func (d *HttpMockDao) FindOneById(id int64) (*model.HttpMock, error) {
	return nil, nil
}

func (d *HttpMockDao) FindByGroup(group sql.NullString) ([]*model.HttpMock, error) {
	return nil, nil
}

func (d *HttpMockDao) FindAll() ([]*model.HttpMock, error) {
	return nil, nil
}
