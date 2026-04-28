package task

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *AppError) Error() string {

	return e.Message
}

var (
	ErrTitleRequired = &AppError{
		Code:    "VALIDATION_ERROR",
		Message: "title is required",
		Status:  400,
	}

	ErrDescriptionRequired = &AppError{
		Code:    "VALIDATION_ERROR",
		Message: "description is required",
		Status:  400,
	}

	ErrAssignedRequired = &AppError{
		Code:    "VALIDATION_ERROR",
		Message: "assigned user is required",
		Status:  400,
	}

	ErrTaskNotFound = &AppError{
		Code:    "VALIDATION_ERROR",
		Message: "empty field",
		Status:  400,
	}

	ErrNothingToUpdate = &AppError{
		Code:    "VALIDATION_ERROR",
		Message: "empty field",
		Status:  400,
	}

	ErrInternal = &AppError{
		Code:    "INTERNAL_ERROR",
		Message: "something went wrong",
		Status:  500,
	}
)
