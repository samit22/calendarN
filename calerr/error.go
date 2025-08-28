package calerr

type NewError struct {
	originalError error
	message       string
	httpStatus    int
}

func New(err error, message string, httpStatus int) *NewError {
	return &NewError{
		originalError: err,
		message:       message,
		httpStatus:    httpStatus,
	}
}

func (n *NewError) Error() string {
	return n.message
}

func (n *NewError) GetOriginalError() error {
	return n.originalError
}

func (n *NewError) GetHTTPStatus() int {
	return n.httpStatus
}
