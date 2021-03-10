package error

type HttpError struct {
	HttpStatusCode int
	Err            Error
}

func (e *HttpError) Error() string {
	return e.Err.Error()
}

func NewHttpError(HttpStatusCode int, ErrCode, Msg string) HttpError {
	e := HttpError{
		HttpStatusCode: HttpStatusCode,
		Err: Error{
			ErrCode: ErrCode,
			Msg:     Msg,
		},
	}
	return e
}
