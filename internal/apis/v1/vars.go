package v1

import "net/http"

var (
	StatusOk            = regSuccessCode(200, http.StatusOK, "OK")
	OperationSuccessful = regSuccessCode(200, http.StatusOK, "Operation completed successfully.")

	ErrBadRequest          = regErrCode(400, http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax.")
	ErrUnauthorized        = regErrCode(401, http.StatusUnauthorized, "Authentication is required. Please ensure you provide valid credentials.")
	ErrNotFound            = regErrCode(404, http.StatusNotFound, "The requested resource could not be found.")
	ErrInternalServerError = regErrCode(500, http.StatusInternalServerError, "An internal server error has occurred.")

	ErrEmailAlreadyInUse  = regErrCode(1001, http.StatusBadRequest, "The provided email address is already in use.")
	ErrUserAlreadyExists  = regErrCode(1003, http.StatusBadRequest, "A user with the specified details already exists.")
	ErrUniversityIdExists = regErrCode(1004, http.StatusBadRequest, "A user with the specified university ID already exists.")
	ErrInvalidStudentId   = regErrCode(1005, http.StatusBadRequest, "The provided student ID is invalid.")
	ErrInvalidPassword    = regErrCode(1006, http.StatusBadRequest, "The provided password does not meet validation criteria.")
	ErrLinkExpired        = regErrCode(1007, http.StatusBadRequest, "The confirmation link has expired.")
	ErrLinkAlreadyUsed    = regErrCode(1008, http.StatusBadRequest, "The confirmation link has already been utilized.")
	ErrEmailNotVerified   = regErrCode(1009, http.StatusForbidden, "Access is forbidden: Email verification is required.")
	ErrPasswordNotMatch   = regErrCode(1010, http.StatusForbidden, "The password entered do not match. Please try again.")
)
