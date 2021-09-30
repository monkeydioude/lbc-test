package response

import (
	"fmt"
	"net/http"
)

// HttpError is a quick wrapper for answering anything but response code 200 :(
func HttpError(rw http.ResponseWriter, msg string, code int) {
	http.Error(rw, fmt.Sprintf(`{"code": %d, "message": "%s"}`, code, msg), code)
}

// NotFound wrapper of a wrapper forcing 404 response code
func NotFound(rw http.ResponseWriter) {
	HttpError(rw, "not found", 404)
}

// BadRequest wrapper of a wrapper forcing 403 response code
func BadRequest(rw http.ResponseWriter) {
	HttpError(rw, "bad request", 403)
}

// InternalServer wrapper of a wrapper forcing 500 response code
func InternalServer(rw http.ResponseWriter) {
	HttpError(rw, "bad request", 500)
}

func Ok(rw http.ResponseWriter, content []byte) {
	rw.Write(content)
}
