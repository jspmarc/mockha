package mapper

import (
	"database/sql"
	"github.com/jspmarc/mockha/dto/http_mock"
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils"
	"regexp"
)

func CreateRequestToModelHttpMock(createRequest *http_mock.CreateRequest) *model.HttpMock {
	var group string
	if createRequest.Group != nil {
		group = *createRequest.Group
	} else {
		group = ""
	}

	isGroupEmpty, _ := regexp.MatchString(`(\s+|)`, group)

	mock := &model.HttpMock{
		Group: sql.NullString{
			String: group,
			Valid:  isGroupEmpty,
		},
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
		RequestHeader:            utils.BytePtrToSqlNullByte(createRequest.RequestHeader),
		RequestBody:              utils.StrPtrToSqlNullString(createRequest.RequestBody),
		RequestBodyMimeType:      utils.StrPtrToSqlNullString(createRequest.RequestBodyMimeType),
		AdditionalResponseHeader: utils.BytePtrToSqlNullByte(createRequest.AdditionalResponseHeader),
		ResponseBody:             utils.StrPtrToSqlNullString(createRequest.ResponseBody),
		ResponseBodyMimeType:     utils.StrPtrToSqlNullString(createRequest.ResponseBodyMimeType),
		ResponseCode:             responseCode,
	}
	return rr
}
