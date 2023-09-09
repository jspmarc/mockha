package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/api/dao"
	"github.com/jspmarc/mockha/model"
)

func NewRequestResponseDao(db *sqlx.DB) dao.HttpRequestResponseDao {
	requestResponseDao := &HttpRequestResponsesDao{}

	requestResponseDao.db = db

	return requestResponseDao
}

type HttpRequestResponsesDao struct {
	db *sqlx.DB
}

func (rr *HttpRequestResponsesDao) Save(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error) {
	return nil, nil
}

func (rr *HttpRequestResponsesDao) Upsert(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error) {
	return nil, nil
}

func (rr *HttpRequestResponsesDao) Delete(id int64) error {
	return nil
}

func (rr *HttpRequestResponsesDao) FindOne(id int64) (*model.HttpRequestResponse, error) {
	return nil, nil
}

func (rr *HttpRequestResponsesDao) FindAllByMockId(mockId int64) ([]*model.HttpRequestResponse, error) {
	return nil, nil
}
