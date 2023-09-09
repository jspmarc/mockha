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

func TestHttpMockDao_Save_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Group:  sql.NullString{String: "group", Valid: true},
		Path:   "path",
		Method: constants.HTTP_METHOD_GET,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO http_mock").
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newMockData, err := daoInstance.Save(httpMockData)
	assert.Nil(t, err)
	assert.Equal(t, httpMockData, newMockData)
}

func TestHttpMockDao_Save_errorBeginTx(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Group:  sql.NullString{String: "group", Valid: true},
		Path:   "path",
		Method: constants.HTTP_METHOD_GET,
	}

	dbError := errors.New("tx")

	mock.ExpectBegin().WillReturnError(dbError)

	newMockData, err := daoInstance.Save(httpMockData)
	assert.Equal(t, err, dbError)
	assert.Nil(t, newMockData)
}

func TestHttpMockDao_Save_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Group:  sql.NullString{String: "group", Valid: true},
		Path:   "path",
		Method: "GET",
	}

	dbError := errors.New("tx")

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO http_mock").
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnError(dbError)
	mock.ExpectCommit()

	newMockData, err := daoInstance.Save(httpMockData)
	assert.Equal(t, err, dbError)
	assert.Nil(t, newMockData)
}

func TestHttpMockDao_Save_errorCommit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Group:  sql.NullString{String: "group", Valid: true},
		Path:   "path",
		Method: "GET",
	}

	dbError := errors.New("tx")

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO http_mock").
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().
		WillReturnError(dbError)

	newMockData, err := daoInstance.Save(httpMockData)
	assert.Equal(t, err, dbError)
	assert.Nil(t, newMockData)
}

func TestHttpMockDao_Update_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Id:     int64(1),
		Group:  sql.NullString{String: "", Valid: false},
		Path:   "path",
		Method: "GET",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE http_mock").
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	updatedMockData, err := daoInstance.Update(httpMockData)
	assert.Nil(t, err)
	assert.Equal(t, httpMockData, updatedMockData)
}

func TestHttpMockDao_Update_errorBeginTx(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Id:     int64(1),
		Group:  sql.NullString{String: "", Valid: false},
		Path:   "path",
		Method: "GET",
	}

	dbError := errors.New("tx")

	mock.ExpectBegin().
		WillReturnError(dbError)

	updatedMockData, err := daoInstance.Update(httpMockData)
	assert.Equal(t, err, dbError)
	assert.Nil(t, updatedMockData)
}

func TestHttpMockDao_Update_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Id:     int64(1),
		Group:  sql.NullString{String: "", Valid: false},
		Path:   "path",
		Method: "GET",
	}

	dbError := errors.New("tx")

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE http_mock").
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnError(dbError)
	mock.ExpectCommit()

	updatedMockData, err := daoInstance.Update(httpMockData)
	assert.Equal(t, err, dbError)
	assert.Nil(t, updatedMockData)
}

func TestHttpMockDao_Update_errorCommit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	httpMockData := &model.HttpMock{
		Id:     int64(1),
		Group:  sql.NullString{String: "", Valid: false},
		Path:   "path",
		Method: "GET",
	}

	dbError := errors.New("tx")

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE http_mock").
		WithArgs(httpMockData.Group, httpMockData.Path, httpMockData.Method, httpMockData.Id).
		WillReturnError(dbError)
	mock.ExpectCommit()

	updatedMockData, err := daoInstance.Update(httpMockData)
	assert.Equal(t, err, dbError)
	assert.Nil(t, updatedMockData)
}

func TestHttpMockDao_DeleteById_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	id := int64(1)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM http_mock").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := daoInstance.DeleteById(id)
	assert.Nil(t, err)
}

func TestHttpMockDao_DeleteById_errorBeginTx(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	dbError := errors.New("tx")

	id := int64(1)

	mock.ExpectBegin().
		WillReturnError(dbError)

	err := daoInstance.DeleteById(id)
	assert.Equal(t, err, dbError)
}

func TestHttpMockDao_DeleteById_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	id := int64(1)

	dbError := errors.New("tx")

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM http_mock").
		WithArgs(id).
		WillReturnError(dbError)
	mock.ExpectCommit()

	err := daoInstance.DeleteById(id)
	assert.Equal(t, err, dbError)
}

func TestHttpMockDao_DeleteById_errorCommit(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	id := int64(1)

	dbError := errors.New("tx")

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM http_mock").
		WithArgs(id).
		WillReturnError(dbError)
	mock.ExpectCommit()

	err := daoInstance.DeleteById(id)
	assert.Equal(t, err, dbError)
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

	mock.ExpectQuery("SELECT (.*) FROM http_mock").
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

	assert.Nil(t, err)
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

	dbError := errors.New("err")

	mock.ExpectQuery("SELECT (.*) FROM http_mock").
		WithArgs(group).
		WillReturnError(dbError)

	actual, err := daoInstance.FindByGroup(group)

	assert.Equal(t, dbError, err)
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

	mock.ExpectQuery("SELECT (.*) FROM http_mock").
		WillReturnRows(rows)

	actual, err := daoInstance.FindAll()

	expected := make([]*model.HttpMock, 1)
	expected[0] = &model.HttpMock{
		Id:     1,
		Group:  group,
		Path:   "",
		Method: constants.HTTP_METHOD_GET,
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestHttpMockDao_FindAll_error(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	daoInstance := dao.NewHttpMockDao(db)

	dbError := errors.New("err")

	mock.ExpectQuery("SELECT (.*) FROM http_mock").
		WillReturnError(dbError)

	actual, err := daoInstance.FindAll()

	assert.Equal(t, dbError, err)
	assert.Nil(t, actual)
}
