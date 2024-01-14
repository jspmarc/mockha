package http_mock

import (
	"github.com/jspmarc/mockha/constants"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils"
)

type CreateRequest struct {
	Group  string
	Path   string
	Method constants.HttpMethod

	RequestHeader       *byte   `json:"requestHeader"`
	RequestBody         *string `json:"requestBody"`
	RequestBodyMimeType *string `json:"requestBodyMimeType"`

	AdditionalResponseHeader *byte   `json:"additionalResponseHeader"`
	ResponseBody             *string `json:"responseBody"`
	ResponseBodyMimeType     *string `json:"responseBodyMimeType"`
	ResponseCode             uint16  `json:"responseCode"`
}

func (cr *CreateRequest) ToModelHttpMock() *model.HttpMock {
	mock := &model.HttpMock{
		Group:  cr.Group,
		Path:   cr.Path,
		Method: cr.Method,
	}
	return mock
}

func (cr *CreateRequest) ToCreateRequestToModelHttpRequestResponse(mockId int64) *model.HttpRequestResponse {
	responseCode := uint16(200)
	if cr.ResponseCode >= 100 && cr.ResponseCode < 600 {
		responseCode = cr.ResponseCode
	}
	rr := &model.HttpRequestResponse{
		HttpMockId:               mockId,
		RequestBody:              utils.StrPtrToSqlNullString(cr.RequestBody),
		RequestBodyMimeType:      utils.StrPtrToSqlNullString(cr.RequestBodyMimeType),
		AdditionalResponseHeader: utils.BytePtrToSqlNullByte(cr.AdditionalResponseHeader),
		ResponseBody:             utils.StrPtrToSqlNullString(cr.ResponseBody),
		ResponseBodyMimeType:     utils.StrPtrToSqlNullString(cr.ResponseBodyMimeType),
		ResponseCode:             responseCode,
	}
	return rr
}
