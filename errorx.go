package errorx

type Errorx struct {
	code    int32
	message string
}

func (x Errorx) Error() string {
	return x.message
}

func (x Errorx) Code() int32 {
	return x.code
}

func NewErrorx(code int32, message string) *Errorx {
	return &Errorx{
		code:    code,
		message: message,
	}
}