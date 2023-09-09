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
	query := rr.db.Rebind("INSERT INTO http_request_response (http_mock_id, request_body, request_body_mime_type, additional_response_header, response_body, response_body_mime_type, response_code) VALUES (?, ?, ?, ?, ?, ?, ?)")

	tx, err := rr.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		query,
		reqres.HttpMockId,
		reqres.RequestBody,
		reqres.RequestBodyMimeType,
		reqres.AdditionalResponseHeader,
		reqres.ResponseBody,
		reqres.ResponseBodyContentType,
		reqres.ResponseCode,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return reqres, nil
}

func (rr *HttpRequestResponsesDao) Update(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error) {
	query := rr.db.Rebind(`
UPDATE http_request_response SET
                                http_mock_id = ?,
                                request_body = ?,
                                request_body_mime_type = ?,
                                additional_response_header = ?,
                                response_body = ?,
                                response_body_mime_type = ?,
                                response_code = ?
WHERE id = ?
`)

	tx, err := rr.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		query,
		reqres.HttpMockId,
		reqres.RequestBody,
		reqres.RequestBodyMimeType,
		reqres.AdditionalResponseHeader,
		reqres.ResponseBody,
		reqres.ResponseBodyContentType,
		reqres.ResponseCode,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return reqres, nil
}

func (rr *HttpRequestResponsesDao) Delete(id int64) error {
	query := rr.db.Rebind("DELETE FROM http_request_response WHERE id = ?")

	tx, err := rr.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (rr *HttpRequestResponsesDao) FindOneById(id int64) (*model.HttpRequestResponse, error) {
	requestResponse := &model.HttpRequestResponse{}

	query := rr.db.Rebind("SELECT * FROM http_request_response WHERE id = ?")
	err := rr.db.Select(&requestResponse, query, id)
	if err != nil {
		return nil, err
	}

	return requestResponse, nil
}

func (rr *HttpRequestResponsesDao) FindOneForRequest(httpMockId int64, reqBody string) (*model.HttpRequestResponse, error) {
	requestResponse := &model.HttpRequestResponse{}

	query := rr.db.Rebind("SELECT * FROM http_request_response WHERE http_mock_id = ? AND request_body = ?")
	err := rr.db.Select(&requestResponse, query, httpMockId, reqBody)
	if err != nil {
		return nil, err
	}

	return requestResponse, nil
}
