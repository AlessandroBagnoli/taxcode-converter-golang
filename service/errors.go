package service

type RuntimeError struct {
	Code    int
	Message string
}

func NewRuntimeError(code int, message string) error {
	return RuntimeError{code, message}
}

func NewValidationError(message string) error {
	return NewRuntimeError(400, message)
}

func NewCityNotPresentError(message string) error {
	return NewRuntimeError(404, message)
}

func (e RuntimeError) Error() string {
	return e.Message
}
