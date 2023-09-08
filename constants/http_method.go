package constants

import "net/http"

type HttpMethod string

const (
	HTTP_METHOD_GET     HttpMethod = http.MethodGet
	HTTP_METHOD_HEAD    HttpMethod = http.MethodHead
	HTTP_METHOD_POST    HttpMethod = http.MethodPost
	HTTP_METHOD_PUT     HttpMethod = http.MethodPut
	HTTP_METHOD_DELETE  HttpMethod = http.MethodDelete
	HTTP_METHOD_CONNECT HttpMethod = http.MethodConnect
	HTTP_METHOD_OPTIONS HttpMethod = http.MethodOptions
	HTTP_METHOD_TRACE   HttpMethod = http.MethodTrace
	HTTP_METHOD_PATCH   HttpMethod = http.MethodPatch
)
