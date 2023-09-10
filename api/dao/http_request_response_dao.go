package dao

import (
	"database/sql"
	"github.com/jspmarc/mockha/model"
)

type HttpRequestResponseDao interface {
	Save(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error)
	Update(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error)
	Delete(id int64) error
	FindOneById(id int64) (*model.HttpRequestResponse, error)
	FindOneForRequest(httpMockId int64, reqBody sql.NullString) (*model.HttpRequestResponse, error)
}
