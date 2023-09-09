package model

import (
	"database/sql"
	"github.com/jspmarc/mockha/constants"
)

type HttpMock struct {
	Id int64

	Group  sql.NullString
	Path   string
	Method constants.HttpMethod
}
