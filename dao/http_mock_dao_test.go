package dao_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/dao"
	"github.com/jspmarc/mockha/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	httpMockData = model.HttpMock{
		Group:  sql.NullString{String: "group", Valid: true},
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

func TestHttpMockDao_Save_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectExec(httpMockInsertQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnResult(sqlmock.NewResult(2, 1))

	expected := httpMockData
	expected.Id = 2

	newMockData, err := daoInstance.Save(&httpMockData)
	assert.NoError(t, err)
	assert.Equal(t, expected, *newMockData)
}
func TestHttpMockDao_Save_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectExec(httpMockInsertQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnError(dbError)

	newMockData, err := daoInstance.Save(&httpMockData)
	assert.Error(t, err, dbError)
	assert.Nil(t, newMockData)
}

func TestHttpMockDao_Update_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectExec(httpMockUpdateQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	updatedMockData, err := daoInstance.Update(&httpMockData)
	assert.NoError(t, err)
	assert.Equal(t, httpMockData, *updatedMockData)
}

func TestHttpMockDao_Update_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectExec(httpMockUpdateQuery).
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnError(dbError)

	updatedMockData, err := daoInstance.Update(&httpMockData)
	assert.Error(t, err, dbError)
	assert.Nil(t, updatedMockData)
}

func TestHttpMockDao_DeleteById_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectExec(httpMockDeleteQuery).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := daoInstance.DeleteById(id)
	assert.NoError(t, err)
}

func TestHttpMockDao_DeleteById_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectExec(httpMockDeleteQuery).
		WithArgs(id).
		WillReturnError(dbError)

	err := daoInstance.DeleteById(id)
	assert.Error(t, err, dbError)
}

func TestHttpMockDao_FindByGroup_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	group := sql.NullString{
		String: "group",
		Valid:  true,
	}

	rows := sqlmock.NewRows([]string{"id", "group", "path", "method"}).
		AddRow(1, group, "", constants.HTTP_METHOD_GET)

	mock.ExpectQuery(httpMockFindByGroupQuery).
		WithArgs(group).
		WillReturnRows(rows)

	actual, err := daoInstance.FindByGroup(group)

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

func TestHttpMockDao_FindByGroup_error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	group := sql.NullString{
		String: "group",
		Valid:  true,
	}

	mock.ExpectQuery(httpMockFindByGroupQuery).
		WithArgs(group).
		WillReturnError(dbError)

	actual, err := daoInstance.FindByGroup(group)

	assert.Error(t, dbError, err)
	assert.Nil(t, actual)
}

func TestHttpMockDao_FindAll_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	group := sql.NullString{
		String: "group",
		Valid:  true,
	}

	rows := sqlmock.NewRows([]string{"id", "group", "path", "method"}).
		AddRow(1, group, "", constants.HTTP_METHOD_GET)

	mock.ExpectQuery(httpMockFindAllQuery).
		WillReturnRows(rows)

	actual, err := daoInstance.FindAll()

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

func TestHttpMockDao_FindAll_error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	mock.ExpectQuery(httpMockFindAllQuery).
		WillReturnError(dbError)

	actual, err := daoInstance.FindAll()

	assert.Error(t, dbError, err)
	assert.Nil(t, actual)
}
