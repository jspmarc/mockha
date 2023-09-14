package dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/api/dao"
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
	query := d.db.Rebind("INSERT INTO http_mock (mock_group, path, method) VALUES (?, ?, ?) RETURNING id")

	result, err := d.db.Exec(query, mock.Group, mock.Path, mock.Method)
	if err != nil {
		return nil, err
	}
	mock.Id, _ = result.LastInsertId()

	return mock, nil
}

func (d *HttpMockDao) Update(mock *model.HttpMock) (*model.HttpMock, error) {
	query := d.db.Rebind("UPDATE http_mock SET method = ?, mock_group = ?, path = ? WHERE id = ?")

	_, err := d.db.Exec(query, mock.Group, mock.Path, mock.Method, mock.Id)
	if err != nil {
		return nil, err
	}

	return mock, nil
}

func (d *HttpMockDao) DeleteById(id int64) error {
	query := d.db.Rebind("DELETE FROM http_mock WHERE id = ?")

	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *HttpMockDao) FindByGroup(group sql.NullString) ([]*model.HttpMock, error) {
	mocks := make([]*model.HttpMock, 0)

	query := d.db.Rebind(`SELECT * FROM http_mock WHERE mock_group = ?`)
	err := d.db.Select(&mocks, query, group)
	if err != nil {
		return nil, err
	}

	return mocks, nil
}

func (d *HttpMockDao) FindAll() ([]*model.HttpMock, error) {
	mocks := make([]*model.HttpMock, 0)

	err := d.db.Select(&mocks, "SELECT * FROM http_mock")
	if err != nil {
		return nil, err
	}

	return mocks, nil
}
