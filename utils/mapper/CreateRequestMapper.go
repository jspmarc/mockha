package mapper

import (
	"github.com/jspmarc/mockha/dto/http_mock"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils"
)

func CreateRequestToModelHttpMock(createRequest *http_mock.CreateRequest) *model.HttpMock {
	mock := &model.HttpMock{
		Group:  createRequest.Group,
		Path:   createRequest.Path,
		Method: createRequest.Method,
	}
	return mock
}

func CreateRequestToModelHttpRequestResponse(createRequest *http_mock.CreateRequest, mockId int64) *model.HttpRequestResponse {
	responseCode := uint16(200)
	if createRequest.ResponseCode >= 100 && createRequest.ResponseCode < 600 {
		responseCode = createRequest.ResponseCode
	}
	rr := &model.HttpRequestResponse{
		HttpMockId:               mockId,
		RequestBody:              utils.StrPtrToSqlNullString(createRequest.RequestBody),
		RequestBodyMimeType:      utils.StrPtrToSqlNullString(createRequest.RequestBodyMimeType),
		AdditionalResponseHeader: utils.BytePtrToSqlNullByte(createRequest.AdditionalResponseHeader),
		ResponseBody:             utils.StrPtrToSqlNullString(createRequest.ResponseBody),
		ResponseBodyMimeType:     utils.StrPtrToSqlNullString(createRequest.ResponseBodyMimeType),
		ResponseCode:             responseCode,
	}
	return rr
}
