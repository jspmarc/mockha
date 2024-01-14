package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	reqres = model.HttpRequestResponse{
		HttpMockId: 1,
		RequestBody: sql.NullString{
			String: "request-body",
			Valid:  true,
		},
		RequestBodyMimeType: sql.NullString{
			String: "text/plain",
			Valid:  true,
		},
		AdditionalResponseHeader: sql.NullByte{
			Valid: false,
		},
		ResponseBody: sql.NullString{
			Valid: false,
		},
		ResponseBodyMimeType: sql.NullString{
			Valid: false,
		},
		ResponseCode: 200,
	}

	httpRequestResponseExpected = &model.HttpRequestResponse{
		Id:         1,
		HttpMockId: 1,

		RequestHeader:       sql.NullByte{},
		RequestBody:         sql.NullString{},
		RequestBodyMimeType: sql.NullString{},

		AdditionalResponseHeader: sql.NullByte{},
		ResponseBody:             sql.NullString{},
		ResponseBodyMimeType:     sql.NullString{},
		ResponseCode:             200,
	}

	httpRequestResponseInsertQuery         = "INSERT INTO http_request_response (.+) VALUES (.+)"
	httpRequestResponseUpdateQuery         = "UPDATE http_request_response SET (.+) WHERE id = (.+)"
	httpRequestResponseDeleteQuery         = "DELETE FROM http_request_response WHERE id = (.+)"
	httpRequestResponseFindByIdQuery       = "SELECT (.*) FROM http_request_response WHERE id = (.+)"
	httpRequestResponseFindForRequestQuery = "SELECT (.*) FROM http_request_response WHERE http_mock_id = (.+) AND request_body = (.)"
)

func TestHttpRequestResponsesRepository_Save_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectExec(httpRequestResponseInsertQuery).
		WithArgs(reqres.HttpMockId, reqres.RequestBody, reqres.RequestBodyMimeType, reqres.AdditionalResponseHeader,
			reqres.ResponseBody, reqres.ResponseBodyMimeType, reqres.ResponseCode).
		WillReturnResult(sqlmock.NewResult(1, 1))

	expected := reqres

	res, err := rrInstance.Save(&reqres)

	assert.NoError(t, err)
	assert.Equal(t, expected, *res)
}

func TestHttpRequestResponsesRepository_Save_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectExec(httpRequestResponseInsertQuery).
		WithArgs(
			reqres.HttpMockId,
			reqres.RequestBody,
			reqres.RequestBodyMimeType,
			reqres.AdditionalResponseHeader,
			reqres.ResponseBody,
			reqres.ResponseBodyMimeType,
			reqres.ResponseCode,
		).
		WillReturnError(dbError)

	res, err := rrInstance.Save(&reqres)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestHttpRequestResponsesRepository_Update_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectExec(httpRequestResponseUpdateQuery).
		WithArgs(reqres.HttpMockId, reqres.RequestBody, reqres.RequestBodyMimeType, reqres.AdditionalResponseHeader,
			reqres.ResponseBody, reqres.ResponseBodyMimeType, reqres.ResponseCode).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := rrInstance.Update(&reqres)

	assert.NoError(t, err)
	assert.Equal(t, reqres, *res)
}

func TestHttpRequestResponsesRepository_Update_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectExec(httpRequestResponseUpdateQuery).
		WithArgs(
			reqres.HttpMockId,
			reqres.RequestBody,
			reqres.RequestBodyMimeType,
			reqres.AdditionalResponseHeader,
			reqres.ResponseBody,
			reqres.ResponseBodyMimeType,
			reqres.ResponseCode,
		).
		WillReturnError(dbError)

	res, err := rrInstance.Update(&reqres)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestHttpRequestResponsesRepository_Delete_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectExec(httpRequestResponseDeleteQuery).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := rrInstance.Delete(id)

	assert.NoError(t, err)
}

func TestHttpRequestResponsesRepository_Delete_errorExec(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectExec(httpRequestResponseDeleteQuery).
		WithArgs(id).
		WillReturnError(dbError)

	err := rrInstance.Delete(id)

	assert.Error(t, err)
}

func TestHttpRequestResponsesRepository_FindOneById_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rows := sqlmock.NewRows([]string{
		"id",
		"http_mock_id",

		"request_header",
		"request_body",
		"request_body_mime_type",

		"additional_response_header",
		"response_body",
		"response_body_mime_type",
		"response_code",
	}).
		AddRow(
			id,
			id,

			sql.NullByte{},
			sql.NullString{},
			sql.NullString{},

			sql.NullByte{},
			sql.NullString{},
			sql.NullString{},
			200,
		)

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectQuery(httpRequestResponseFindByIdQuery).
		WithArgs(id).
		WillReturnRows(rows)

	// call the 'Delete' method
	actual, err := rrInstance.FindOneById(id)

	assert.NoError(t, err)
	assert.Equal(t, httpRequestResponseExpected, actual)
}

func TestHttpRequestResponsesRepository_FindOneById_errorGet(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectQuery(httpRequestResponseFindByIdQuery).
		WithArgs(id).
		WillReturnError(dbError)

	// call the 'Delete' method
	actual, err := rrInstance.FindOneById(id)

	assert.Error(t, err)
	assert.Nil(t, actual)
}

func TestHttpRequestResponsesRepository_FindOneForRequest_success(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rows := sqlmock.NewRows([]string{
		"id",
		"http_mock_id",

		"request_header",
		"request_body",
		"request_body_mime_type",

		"additional_response_header",
		"response_body",
		"response_body_mime_type",
		"response_code",
	}).
		AddRow(
			id,
			id,

			sql.NullByte{},
			sql.NullString{},
			sql.NullString{},

			sql.NullByte{},
			sql.NullString{},
			sql.NullString{},
			200,
		)

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectQuery(httpRequestResponseFindForRequestQuery).
		WithArgs(id, sql.NullString{}).
		WillReturnRows(rows)

	// call the 'Delete' method
	actual, err := rrInstance.FindOneForRequest(id, sql.NullString{})

	assert.NoError(t, err)
	assert.Equal(t, httpRequestResponseExpected, actual)
}

func TestHttpRequestResponsesRepository_FindOneForRequest_errorGet(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	db := sqlx.NewDb(mockDb, "sqlmock")

	rrInstance := repository.NewRequestResponseRepository(db)

	mock.ExpectQuery(httpRequestResponseFindForRequestQuery).
		WithArgs(id, sql.NullString{}).
		WillReturnError(dbError)

	// call the 'Delete' method
	actual, err := rrInstance.FindOneForRequest(id, sql.NullString{})

	assert.Error(t, err)
	assert.Nil(t, actual)
}
