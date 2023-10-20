package model

import (
	"database/sql"
)

type HttpRequestResponse struct {
	Id int64

	HttpMockId int64 `db:"http_mock_id" json:"httpMockId"`

	RequestHeader       sql.NullByte   `db:"request_header" json:"requestHeader"`
	RequestBody         sql.NullString `db:"request_body" json:"requestBody"`
	RequestBodyMimeType sql.NullString `db:"request_body_mime_type" json:"requestBodyMimeType"`

	AdditionalResponseHeader sql.NullByte   `db:"additional_response_header" json:"additionalResponseHeader"`
	ResponseBody             sql.NullString `db:"response_body" json:"responseBody"`
	ResponseBodyMimeType     sql.NullString `db:"response_body_mime_type" json:"responseBodyMimeType"`
	ResponseCode             uint16         `db:"response_code" json:"responseCode"`
}
