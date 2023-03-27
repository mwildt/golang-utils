package errorutils

type NotFound struct {
	Message string
}

func (err NotFound) Error() string {
	return err.Message
}

type ClientError struct {
	Message string
}

func (err ClientError) Error() string {
	return err.Message
}
