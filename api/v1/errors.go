package v1

var (
	ErrSuccess             = newError(200, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	ErrEmailAlreadyUse    = newError(1001, "The email is already in use.")
	ErrUserAlreadyExists  = newError(1003, "User already exists.")
	ErrUniIdAlreadyExists = newError(1004, "User with university id already exists.")
	ErrInvalidStudentId   = newError(1005, "The student is not valid.")
	ErrInvalidPassword    = newError(1006, "The password is not valid.")
	ErrLinkExpired        = newError(1007, "The confirmation link is expired.")
	ErrLinkAlreadyUsed    = newError(1008, "Link already used.")
	ErrEmailNotVerified   = newError(1009, "Access Forbidden: Email verification is required.")
)
