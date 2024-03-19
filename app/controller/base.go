package controller

type Base struct{}

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (b *Base) res(code int, message string, data interface{}) Res {
	return Res{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (b *Base) success(data interface{}) Res {
	return b.res(0, "返回成功", data)
}

func (b *Base) error(message string) Res {
	return b.res(1, message, nil)
}
