package service

type RuntimeError struct {
	Code    int
	Message string
}

func NewRuntimeError(code int, message string) error {
	return RuntimeError{code, message}
}

func (e RuntimeError) Error() string {
	return e.Message
}
