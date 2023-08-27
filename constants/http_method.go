package constants

type HttpMethod string

const (
	HTTP_METHOD_GET     HttpMethod = "GET"
	HTTP_METHOD_HEAD    HttpMethod = "HEAD"
	HTTP_METHOD_POST    HttpMethod = "POST"
	HTTP_METHOD_PUT     HttpMethod = "PUT"
	HTTP_METHOD_DELETE  HttpMethod = "DELETE"
	HTTP_METHOD_CONNECT HttpMethod = "CONNECT"
	HTTP_METHOD_OPTIONS HttpMethod = "OPTIONS"
	HTTP_METHOD_TRACE   HttpMethod = "TRACE"
	HTTP_METHOD_PATCH   HttpMethod = "PATCH"
)
