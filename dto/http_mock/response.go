package http_mock

import (
	"github.com/jspmarc/mockha/model"
)

type Response struct {
	*model.HttpMock
	RequestResponseCombinations []*httpRequestResponse `json:"RequestResponseCombinations"`
}

func NewHttpMockResponse(httpMock *model.HttpMock, requestResponseCombinations ...*model.HttpRequestResponse) *Response {
	resp := &Response{HttpMock: httpMock, RequestResponseCombinations: nil}

	if len(requestResponseCombinations) > 0 {
		resp.RequestResponseCombinations = make([]*httpRequestResponse,
			len(requestResponseCombinations))
		for i, combo := range requestResponseCombinations {
			resp.RequestResponseCombinations[i] = httpRequestResponseDtoFromModel(combo)
		}
	}

	return resp
}
