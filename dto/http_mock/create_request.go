package http_mock

import (
	"github.com/jspmarc/mockha/constants"
)

type CreateRequest struct {
	Group  *string
	Path   string
	Method constants.HttpMethod

	RequestBody         *string `json:"request_body"`
	RequestBodyMimeType *string `json:"request_body_mime_type"`

	AdditionalResponseHeader *byte   `json:"additional_response_header"`
	ResponseBody             *string `json:"response_body"`
	ResponseBodyMimeType     *string `json:"response_body_mime_type"`
	ResponseCode             uint16  `json:"response_code"`
}
