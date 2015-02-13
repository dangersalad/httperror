package httperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPError struct {
	error
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (err *HTTPError) Respond(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(err.Code)
	json.NewEncoder(res).Encode(err)
}

func New(code int, messageArgs ...interface{}) *HTTPError {
	var message string
	var err error
	status := http.StatusText(code)
	if len(messageArgs) == 0 {
		message = status
		err = fmt.Errorf("[%d] %s", code, status)
	} else {
		fmtStr := fmt.Sprintf("%v", messageArgs[0])
		message = fmt.Sprintf(fmtStr, messageArgs[1:]...)
		err = fmt.Errorf("[%d] %s - %s", code, status, message)
	}
	return &HTTPError{err, code, status, message}
}

// Convenience methods for 400 errors

func BadRequest(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusBadRequest, messageArgs...)
}

func Unauthorized(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusUnauthorized, messageArgs...)
}

func PaymentRequired(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusPaymentRequired, messageArgs...)
}

func Forbidden(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusForbidden, messageArgs...)
}

func NotFound(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusNotFound, messageArgs...)
}

func MethodNotAllowed(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusMethodNotAllowed, messageArgs...)
}

func NotAcceptable(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusNotAcceptable, messageArgs...)
}

func ProxyAuthRequired(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusProxyAuthRequired, messageArgs...)
}

func RequestTimeout(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusRequestTimeout, messageArgs...)
}

func Conflict(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusConflict, messageArgs...)
}

func Gone(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusGone, messageArgs...)
}

func LengthRequired(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusLengthRequired, messageArgs...)
}

func PreconditionFailed(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusPreconditionFailed, messageArgs...)
}

func RequestEntityTooLarge(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusRequestEntityTooLarge, messageArgs...)
}

func RequestURITooLong(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusRequestURITooLong, messageArgs...)
}

func UnsupportedMediaType(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusUnsupportedMediaType, messageArgs...)
}

func RequestedRangeNotSatisfiable(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusRequestedRangeNotSatisfiable, messageArgs...)
}

func ExpectationFailed(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusExpectationFailed, messageArgs...)
}

func Teapot(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusTeapot, messageArgs...)
}

// Convenience methods for 500 errors

func InternalServerError(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusInternalServerError, messageArgs...)
}

func NotImplemented(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusNotImplemented, messageArgs...)
}

func BadGateway(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusBadGateway, messageArgs...)
}

func ServiceUnavailable(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusServiceUnavailable, messageArgs...)
}

func GatewayTimeout(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusGatewayTimeout, messageArgs...)
}

func HTTPVersionNotSupported(messageArgs ...interface{}) *HTTPError {
	return New(http.StatusHTTPVersionNotSupported, messageArgs...)
}
