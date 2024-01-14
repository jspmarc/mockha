package repository_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	httpMockData = model.HttpMock{
		Group:  "group",
		Path:   "path",
		Method: constants.HTTP_METHOD_GET,
	}

	dbError = errors.New("dbError")

	id = int64(1)

	httpMockInsertQuery      = "INSERT INTO http_mock (.+) VALUES (.+) RETURNING id"
	httpMockUpdateQuery      = "UPDATE http_mock SET (.+) WHERE id = (.+)"
	httpMockDeleteQuery      = "DELETE FROM http_mock WHERE id = (.+)"
	httpMockFindByGroupQuery = "SELECT (.*) FROM http_mock WHERE mock_group = (.+)"
	httpMockFindAllQuery     = "SELECT (.*) FROM http_mock"
)

func TestHttpMockRepository_Save_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectExec(httpMockInsertQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnResult(sqlmock.NewResult(2, 1))

	expected := httpMockData
	expected.Id = 2

	newMockData, err := RepositoryInstance.Save(&httpMockData)
	assert.NoError(t, err)
	assert.Equal(t, expected, *newMockData)
}
func TestHttpMockRepository_Save_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectExec(httpMockInsertQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnError(dbError)

	newMockData, err := RepositoryInstance.Save(&httpMockData)
	assert.Error(t, err, dbError)
	assert.Nil(t, newMockData)
}

func TestHttpMockRepository_Update_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectExec(httpMockUpdateQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	updatedMockData, err := RepositoryInstance.Update(&httpMockData)
	assert.NoError(t, err)
	assert.Equal(t, httpMockData, *updatedMockData)
}

func TestHttpMockRepository_Update_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectExec(httpMockUpdateQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnError(dbError)

	updatedMockData, err := RepositoryInstance.Update(&httpMockData)
	assert.Error(t, err, dbError)
	assert.Nil(t, updatedMockData)
}

func TestHttpMockRepository_DeleteById_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectExec(httpMockDeleteQuery).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := RepositoryInstance.DeleteById(id)
	assert.NoError(t, err)
}

func TestHttpMockRepository_DeleteById_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectExec(httpMockDeleteQuery).
		WithArgs(id).
		WillReturnError(dbError)

	err := RepositoryInstance.DeleteById(id)
	assert.Error(t, err, dbError)
}

func TestHttpMockRepository_FindByGroup_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	group := "group"

	rows := sqlmock.NewRows([]string{"id", "group", "path", "method"}).
		AddRow(1, group, "", constants.HTTP_METHOD_GET)

	mock.ExpectQuery(httpMockFindByGroupQuery).
		WithArgs(group).
		WillReturnRows(rows)

	actual, err := RepositoryInstance.FindByGroup(group)

	expected := make([]*model.HttpMock, 1)
	expected[0] = &model.HttpMock{
		Id:     1,
		Group:  group,
		Path:   "",
		Method: constants.HTTP_METHOD_GET,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestHttpMockRepository_FindByGroup_error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	group := "group"

	mock.ExpectQuery(httpMockFindByGroupQuery).
		WithArgs(group).
		WillReturnError(dbError)

	actual, err := RepositoryInstance.FindByGroup(group)

	assert.Error(t, dbError, err)
	assert.Nil(t, actual)
}

func TestHttpMockRepository_FindAll_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	group := "group"

	rows := sqlmock.NewRows([]string{"id", "group", "path", "method"}).
		AddRow(1, group, "", constants.HTTP_METHOD_GET)

	mock.ExpectQuery(httpMockFindAllQuery).
		WillReturnRows(rows)

	actual, err := RepositoryInstance.FindAll()

	expected := make([]*model.HttpMock, 1)
	expected[0] = &model.HttpMock{
		Id:     1,
		Group:  group,
		Path:   "",
		Method: constants.HTTP_METHOD_GET,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestHttpMockRepository_FindAll_error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	RepositoryInstance := repository.NewHttpMockRepository(db)

	mock.ExpectQuery(httpMockFindAllQuery).
		WillReturnError(dbError)

	actual, err := RepositoryInstance.FindAll()

	assert.Error(t, dbError, err)
	assert.Nil(t, actual)
}
