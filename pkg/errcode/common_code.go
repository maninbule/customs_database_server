// 编写各种错误代号以及解释

package errcode

var (
	Success                       = NewError(0, "成功")
	ServerError                   = NewError(1000, "服务器内部错误")
	InvalidParams                 = NewError(1001, "入参错误")
	NotFound                      = NewError(1002, "找不到")
	UnauthorizedAuthNotExist      = NewError(1003, "授权失败,找不到对应的AppKey和AppSecret")
	UnauthorizedAuthTokenError    = NewError(1004, "授权失败,Token错误")
	UnauthorizedAuthTokenTimeout  = NewError(1005, "授权失败,Token超时")
	UnauthorizedAuthTokenGenerate = NewError(1006, "授权失败,Token生成失败")
	TooManyRequests               = NewError(1007, "请求过多")
)
