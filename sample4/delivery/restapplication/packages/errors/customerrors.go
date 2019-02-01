package customerrors

// New returns an error that formats as the given text.
func BadRequest(message string) error {
	return &CustomError{400, message, false}
}

func InternalError(message string) error {
	return &CustomError{500, message, false}
}

func NotFoundError(message string) error {
	return &CustomError{404, message, false}
}

func PreConditionFailedError(message string) error {
	return &CustomError{412, message, false}
}

func UnAuthorizedError(message string) error {
	return &CustomError{401, message, false}
}

func UnprocessableEntityError(message string) error {
	return &CustomError{422, message, false}
}

// errorString is a trivial implementation of error.
type CustomError struct {
	statusCode int
	message    string
	status     bool
}

func (e *CustomError) Error() string {
	return e.message
}

func (e *CustomError) GetStatusCode() int {
	return e.statusCode
}

func (e *CustomError) GetStatus() bool {
	return e.status
}

func (e *CustomError) GetMessage() string {
	return e.message
}
