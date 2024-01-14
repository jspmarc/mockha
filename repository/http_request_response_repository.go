package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/jspmarc/mockha/api/repository"
	"github.com/jspmarc/mockha/model"
)

func NewRequestResponseRepository(db *sqlx.DB) repository.HttpRequestResponseRepository {
	requestResponseRepository := &HttpRequestResponsesRepository{}

	requestResponseRepository.db = db

	return requestResponseRepository
}

type HttpRequestResponsesRepository struct {
	db *sqlx.DB
}

func (rr *HttpRequestResponsesRepository) Save(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error) {
	query := rr.db.Rebind("INSERT INTO http_request_response (http_mock_id, request_body, request_body_mime_type, additional_response_header, response_body, response_body_mime_type, response_code) VALUES (?, ?, ?, ?, ?, ?, ?)")

	_, err := rr.db.Exec(
		query,
		reqres.HttpMockId,
		reqres.RequestBody,
		reqres.RequestBodyMimeType,
		reqres.AdditionalResponseHeader,
		reqres.ResponseBody,
		reqres.ResponseBodyMimeType,
		reqres.ResponseCode,
	)
	if err != nil {
		return nil, err
	}

	return reqres, nil
}

func (rr *HttpRequestResponsesRepository) Update(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error) {
	query := rr.db.Rebind(`UPDATE http_request_response SET
                                http_mock_id = ?,
                                request_body = ?,
                                request_body_mime_type = ?,
                                additional_response_header = ?,
                                response_body = ?,
                                response_body_mime_type = ?,
                                response_code = ?
WHERE id = ?`)

	_, err := rr.db.Exec(
		query,
		reqres.HttpMockId,
		reqres.RequestBody,
		reqres.RequestBodyMimeType,
		reqres.AdditionalResponseHeader,
		reqres.ResponseBody,
		reqres.ResponseBodyMimeType,
		reqres.ResponseCode,
	)
	if err != nil {
		return nil, err
	}

	return reqres, nil
}

func (rr *HttpRequestResponsesRepository) Delete(id int64) error {
	query := rr.db.Rebind("DELETE FROM http_request_response WHERE id = ?")

	_, err := rr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (rr *HttpRequestResponsesRepository) FindOneById(id int64) (*model.HttpRequestResponse, error) {
	requestResponse := &model.HttpRequestResponse{}

	query := rr.db.Rebind("SELECT * FROM http_request_response WHERE id = ?")
	err := rr.db.Get(requestResponse, query, id)
	if err != nil {
		return nil, err
	}

	return requestResponse, nil
}

func (rr *HttpRequestResponsesRepository) FindOneForRequest(httpMockId int64, reqBody sql.NullString) (*model.HttpRequestResponse, error) {
	requestResponse := &model.HttpRequestResponse{}

	query := rr.db.Rebind("SELECT * FROM http_request_response WHERE http_mock_id = ? AND request_body = ?")
	err := rr.db.Get(requestResponse, query, httpMockId, reqBody)
	if err != nil {
		return nil, err
	}

	return requestResponse, nil
}
