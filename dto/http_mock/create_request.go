package http_mock

import (
	"github.com/jspmarc/mockha/constants"
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
