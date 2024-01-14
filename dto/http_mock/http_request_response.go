package http_mock

import (
	"github.com/jspmarc/mockha/model"
	"github.com/jspmarc/mockha/utils"
)

type httpRequestResponse struct {
	Id int64 `json:"id"`

	HttpMockId int64 `json:"httpMockId"`

	RequestHeader       *byte   `json:"requestHeader"`
	RequestBody         *string `json:"requestBody"`
	RequestBodyMimeType *string `json:"requestBodyMimeType"`

	AdditionalResponseHeader *byte   `json:"additionalResponseHeader"`
	ResponseBody             *string `json:"responseBody"`
	ResponseBodyMimeType     *string `json:"responseBodyMimeType"`
	ResponseCode             uint16  `json:"responseCode"`
}

func httpRequestResponseDtoFromModel(model *model.HttpRequestResponse) *httpRequestResponse {
	return &httpRequestResponse{
		Id:                       model.Id,
		HttpMockId:               model.HttpMockId,
		RequestHeader:            utils.SqlNullByteToBytePtr(model.RequestHeader),
		RequestBody:              utils.SqlNullStringToStrPtr(model.RequestBody),
		RequestBodyMimeType:      utils.SqlNullStringToStrPtr(model.RequestBody),
		AdditionalResponseHeader: utils.SqlNullByteToBytePtr(model.AdditionalResponseHeader),
		ResponseBody:             utils.SqlNullStringToStrPtr(model.ResponseBody),
		ResponseBodyMimeType:     utils.SqlNullStringToStrPtr(model.ResponseBodyMimeType),
		ResponseCode:             model.ResponseCode,
	}
}
