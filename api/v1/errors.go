package v1

var (
	// common errors
	ErrSuccess             = newError(200, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrEmailAlreadyUse    = newError(1001, "The email is already in use.")
	ErrUserAlreadyExists  = newError(400, "User already exists.")
	ErrUniIdAlreadyExists = newError(400, "User with university id already exists.")
	ErrInvalidStudentId   = newError(400, "The student is not valid.")
	ErrInvalidPassword    = newError(400, "The password is not valid.")
	ErrLinkExpired        = newError(400, "The confirmation link is expired.")
	ErrLinkAlreadyUsed    = newError(400, "Link already used.")
	ErrEmailNotVerified   = newError(403, "Access Forbidden: Email verification is required.")
)
