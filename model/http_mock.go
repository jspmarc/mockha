package model

import (
	"database/sql"
	"github.com/josep/mockha/utils"
)

type HttpMock struct {
	Group  sql.NullString
	Path   string
	Method utils.HttpMethod

	RequestHeader          sql.NullByte   `db:"request_header" json:"request_header"`
	RequestBody            sql.NullString `db:"request_body" json:"request_body"`
	RequestBodyContentType sql.NullString `db:"request_body_content_type" json:"request_body_content_type"`

	AdditionalResponseHeader sql.NullByte   `db:"additional_response_header" json:"additional_response_header"`
	ResponseBody             sql.NullString `db:"response_body" json:"response_body"`
	ResponseBodyContentType  sql.NullString `db:"response_body_content_type" json:"response_body_content_type"`
	ResponseCode             uint16         `db:"response_code" json:"response_code"`
}
