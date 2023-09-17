package model

import (
	"database/sql"
)

type HttpRequestResponse struct {
	Id int64

	HttpMockId int64 `db:"http_mock_id"`

	RequestBody         sql.NullString `db:"request_body"`
	RequestBodyMimeType sql.NullString `db:"request_body_mime_type"`

	AdditionalResponseHeader sql.NullByte   `db:"additional_response_header"`
	ResponseBody             sql.NullString `db:"response_body"`
	ResponseBodyMimeType     sql.NullString `db:"response_body_mime_type"`
	ResponseCode             uint16         `db:"response_code"`
}
