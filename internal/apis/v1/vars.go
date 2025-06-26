package v1

import "net/http"

var (
	StatusOk            = regSuccessCode(200, http.StatusOK, "OK")
	OperationSuccessful = regSuccessCode(200, http.StatusOK, "Operation completed successfully.")

	ErrBadRequest          = regErrCode(400, http.StatusBadRequest, "The request could not be understood by the server due to malformed syntax.")
	ErrUnauthorized        = regErrCode(401, http.StatusUnauthorized, "Authentication is required. Please ensure you provide valid credentials.")
	ErrForbiddenAction     = regErrCode(403, http.StatusForbidden, "You do not have permission to do this action.")
	ErrNotFound            = regErrCode(404, http.StatusNotFound, "The requested resource could not be found.")
	ErrUserNotExists       = regErrCode(404, http.StatusNotFound, "The user does not exist.")
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

	ErrBatchAlreadyExists     = regErrCode(1011, http.StatusBadRequest, "A batch with the specified name already exists.")
	ErrBatchNotFound          = regErrCode(1012, http.StatusNotFound, "The specified batch could not be found.")
	ErrStudentNotEligible     = regErrCode(1013, http.StatusBadRequest, "The student is not eligible to register for this batch.")
	ErrInvalidTeamRegDeadline = regErrCode(1013, http.StatusBadRequest, "Team registration deadline should be before pre-defence date.")
	ErrInvalidPreDefenceDate  = regErrCode(1014, http.StatusBadRequest, "Pre-defence date should be before defence date.")
	ErrAlreadyRegForBatch     = regErrCode(1015, http.StatusBadRequest, "The student is already registered for this batch.")
	ErrTeamRegDeadlinePassed  = regErrCode(1016, http.StatusBadRequest, "Team registration deadline has passed. Please contact the admin or convenor for further assistance.")
	ErrBatchCanNotBeClosed    = regErrCode(1017, http.StatusBadRequest, "The specified batch could not be closed.")

	ErrNotBatchRegisterer            = regErrCode(1018, http.StatusForbidden, "User must be registered to the batch in order to create a team.")
	ErrTeamInvitationToNonRegisterer = regErrCode(1019, http.StatusBadRequest, "Target student must be registered to the batch in order to be invited to the team.")
)
