package fault

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrBadRequest = errors.New("bad request")
var ErrUnauthorized = errors.New("unauthorized")
var ErrPermissionDenied = errors.New("permission denied")
var ErrNotFound = errors.New("not found")
var ErrTooManyRequests = errors.New("too many requests")
var ErrEmailExists = errors.New("email exists")
var ErrUndefined = errors.New("undefined error")
var ErrVerificationCode = errors.New("verification code error")
var ErrCreateTableConflict = errors.New("table label or position conflict error")
var ErrItemAttributesConflict = errors.New("attributes conflict error")
var ErrEnum = errors.New("enum parse error")

type ErrorCode int

const (
	BAD_REQUEST_ERROR             ErrorCode = http.StatusBadRequest
	UNATHORIZED_ERROR             ErrorCode = 401
	PERMISSION_ERROR              ErrorCode = 403
	NOT_FOUND_ERROR               ErrorCode = 404
	TOO_MANY_REQUESTS             ErrorCode = 429
	EMAIL_EXISTS_ERROR            ErrorCode = 6002
	AUTHENTICATION_ERROR          ErrorCode = 6003
	SERVICE_KEY_DUPLICATION_ERROR ErrorCode = 6004
	VERIFICATION_CODE_ERROR       ErrorCode = 6005
	TABLE_CONFLICT_ERROR          ErrorCode = 6006
	VERIFICATION_CODE_FREQUENT    ErrorCode = 6007
	ATTRIBUTES_CONFLICT           ErrorCode = 6008
	UNDEFINE_ERROR                ErrorCode = 9999
)

func GetCode(err error) (statusCode int, errorCode ErrorCode) {
	switch {
	case errors.Is(err, ErrBadRequest):
		statusCode = http.StatusBadRequest
		errorCode = BAD_REQUEST_ERROR
	case errors.Is(err, ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		errorCode = UNATHORIZED_ERROR
	case errors.Is(err, ErrPermissionDenied):
		statusCode = http.StatusForbidden
		errorCode = PERMISSION_ERROR
	case errors.Is(err, ErrNotFound):
		statusCode = http.StatusNotFound
		errorCode = NOT_FOUND_ERROR
	case errors.Is(err, ErrTooManyRequests):
		statusCode = http.StatusTooManyRequests
		errorCode = TOO_MANY_REQUESTS
	case errors.Is(err, ErrEmailExists):
		statusCode = http.StatusForbidden
		errorCode = EMAIL_EXISTS_ERROR
	case errors.Is(err, ErrVerificationCode):
		statusCode = http.StatusPreconditionFailed
		errorCode = VERIFICATION_CODE_ERROR
	case errors.Is(err, ErrUndefined):
		statusCode = http.StatusInternalServerError
		errorCode = UNDEFINE_ERROR
	case errors.Is(err, ErrCreateTableConflict):
		statusCode = http.StatusConflict
		errorCode = TABLE_CONFLICT_ERROR
	case errors.Is(err, ErrItemAttributesConflict):
		statusCode = http.StatusConflict
		errorCode = ATTRIBUTES_CONFLICT
	default:
		statusCode = http.StatusInternalServerError
		errorCode = UNDEFINE_ERROR
	}
	return statusCode, errorCode
}

func GinHandler(ctx *gin.Context, err error) {
	statusCode, errorCode := GetCode(err)
	ctx.JSON(statusCode, gin.H{
		"code":    errorCode,
		"message": err.Error(),
	})
}
