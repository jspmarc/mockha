package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/api/repository"
	"github.com/jspmarc/mockha/model"
)

func NewHttpMockRepository(db *sqlx.DB) repository.HttpMockRepository {
	mockRepository := &HttpMockRepository{}

	mockRepository.db = db

	return mockRepository
}

type HttpMockRepository struct {
	db *sqlx.DB
}

func (d *HttpMockRepository) Save(mock *model.HttpMock) (*model.HttpMock, error) {
	query := d.db.Rebind("INSERT INTO http_mock (mock_group, path, method) VALUES (?, ?, ?) RETURNING id")

	result, err := d.db.Exec(query, mock.Group, mock.Path, mock.Method)
	if err != nil {
		return nil, err
	}
	mock.Id, _ = result.LastInsertId()

	return mock, nil
}

func (d *HttpMockRepository) Update(mock *model.HttpMock) (*model.HttpMock, error) {
	query := d.db.Rebind("UPDATE http_mock SET method = ?, mock_group = ?, path = ? WHERE id = ?")

	_, err := d.db.Exec(query, mock.Group, mock.Path, mock.Method, mock.Id)
	if err != nil {
		return nil, err
	}

	return mock, nil
}

func (d *HttpMockRepository) DeleteById(id int64) error {
	query := d.db.Rebind("DELETE FROM http_mock WHERE id = ?")

	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (d *HttpMockRepository) FindByGroup(group string) ([]*model.HttpMock, error) {
	mocks := make([]*model.HttpMock, 0)

	query := d.db.Rebind(`SELECT * FROM http_mock WHERE mock_group = ?`)
	err := d.db.Select(&mocks, query, group)
	if err != nil {
		return nil, err
	}

	return mocks, nil
}

func (d *HttpMockRepository) FindAll() ([]*model.HttpMock, error) {
	mocks := make([]*model.HttpMock, 0)

	err := d.db.Select(&mocks, "SELECT * FROM http_mock")
	if err != nil {
		return nil, err
	}

	return mocks, nil
}
