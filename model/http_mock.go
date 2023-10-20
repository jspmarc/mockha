package model

import (
	"github.com/jspmarc/mockha/constants"
)

type HttpMock struct {
	Id int64

	Group  string               `json:"group"`
	Path   string               `json:"path"`
	Method constants.HttpMethod `json:"method"`
}
