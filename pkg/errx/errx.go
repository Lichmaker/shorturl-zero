package errx



import "fmt"

type CustomError struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func New(c int64, msg string) error {
	return &CustomError{
		Code: c,
		Msg:  msg,
	}
}

func NewWithCode(c int64) error {
	return &CustomError{
		Code: c,
		Msg:  getMsgByCode(c),
	}
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("ErrCode:%d , ErrMsg:%s", c.Code, c.Msg)
}
