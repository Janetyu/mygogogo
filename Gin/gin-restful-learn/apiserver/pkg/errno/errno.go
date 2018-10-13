package errno

import (
	"fmt"
)

type Errno struct {
	Code    int
	Message string
}

// Err represents an error 一个Err代表一个error
type Err struct {
	Code    int
	Message string
	Err     error
}

// 返回错误的信息
func (err *Errno) Error() string {
	return err.Message
}

// 新建定制错误
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFpund.Code
}

// 解析定制错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	// 使用断言进行类型判断
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:

	}

	return InternalServerError.Code, err.Error()

}
