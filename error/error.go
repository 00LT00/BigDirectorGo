package error

import (
	"fmt"
	"runtime"
)

type Error struct {
	ErrCode string
	Msg     string
}

var tag = 1

func caller() string {
	pc, _, _, _ := runtime.Caller(4)
	return runtime.FuncForPC(pc).Name()
}

func (e *Error) Error() string {
	return fmt.Sprintf("发生在: %s, 错误是: %v", caller(), e.Msg)
}

func NewError(Msg string, ErrCode string) Error {
	if ErrCode == "" {
		ErrCode = fmt.Sprintf("%03d", tag)
		tag++
	}
	e := Error{ErrCode: ErrCode, Msg: Msg}
	return e
}
