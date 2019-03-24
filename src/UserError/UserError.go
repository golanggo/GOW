package UserError

type UError struct {
	code int
	msg  string
}

func (e *UError) Error() string {
	return e.msg
}
