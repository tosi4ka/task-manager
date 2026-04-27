package auth

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *AppError) Error() string {

	return e.Message
}

var (
	ErrEmailExists = &AppError{
		Code:    "EMAIL_EXISTS",
		Message: "email already exists",
		Status:  400,
	}

	ErrInvalidCredentials = &AppError{
		Code:    "INVALID_CREDENTIALS",
		Message: "wrong email or password",
		Status:  400,
	}

	ErrUserNotFound = &AppError{
		Code:    "USER_NOT_FOUND",
		Message: "wrong email or password",
		Status:  400,
	}
)
