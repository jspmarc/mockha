package dao

import "github.com/jspmarc/mockha/model"

type HttpRequestResponseDao interface {
	Save(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error)
	Upsert(reqres *model.HttpRequestResponse) (*model.HttpRequestResponse, error)
	Delete(id int64) error
	FindOne(id int64) (*model.HttpRequestResponse, error)
	FindAllByMockId(mockId int64) ([]*model.HttpRequestResponse, error)
}
