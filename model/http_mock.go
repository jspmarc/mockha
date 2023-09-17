package model

import (
	"database/sql"
	"github.com/jspmarc/mockha/constants"
)

type HttpMock struct {
	Id int64

	Group  sql.NullString       `json:"group"`
	Path   string               `json:"path"`
	Method constants.HttpMethod `json:"method"`
}
