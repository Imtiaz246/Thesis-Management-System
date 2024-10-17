package v1

import "net/http"

var (
	ErrSuccess             = registerCodes(200, http.StatusOK, "Operation completed successfully.")
	ErrBadRequest          = registerCodes(400, http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax.")
	ErrUnauthorized        = registerCodes(401, http.StatusUnauthorized, "Authentication is required and has failed or has not yet been provided.")
	ErrNotFound            = registerCodes(404, http.StatusNotFound, "The requested resource could not be found.")
	ErrInternalServerError = registerCodes(500, http.StatusInternalServerError, "An internal server error has occurred.")

	ErrEmailAlreadyInUse  = registerCodes(1001, http.StatusBadRequest, "The provided email address is already in use.")
	ErrUserAlreadyExists  = registerCodes(1003, http.StatusBadRequest, "A user with the specified details already exists.")
	ErrUniversityIdExists = registerCodes(1004, http.StatusBadRequest, "A user with the specified university ID already exists.")
	ErrInvalidStudentId   = registerCodes(1005, http.StatusBadRequest, "The provided student ID is invalid.")
	ErrInvalidPassword    = registerCodes(1006, http.StatusBadRequest, "The provided password does not meet validation criteria.")
	ErrLinkExpired        = registerCodes(1007, http.StatusBadRequest, "The confirmation link has expired.")
	ErrLinkAlreadyUsed    = registerCodes(1008, http.StatusBadRequest, "The confirmation link has already been utilized.")
	ErrEmailNotVerified   = registerCodes(1009, http.StatusForbidden, "Access is forbidden: Email verification is required.")
)
