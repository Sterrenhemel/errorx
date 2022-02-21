package errorx

type Errorx interface {
	Error() string
	Code() int32
}

type Err struct {
	code    int32
	message string
}

func (x Err) Error() string {
	return x.message
}

func (x Err) Code() int32 {
	return x.code
}

func NewErrorx(code int32, message string) Errorx {
	return Err{
		code:    code,
		message: message,
	}
}
