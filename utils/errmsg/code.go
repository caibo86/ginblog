package errmsg

type Code int

const (
	ErrServer          Code = 1
	ErrBadRequest      Code = 2
	ErrUnauthorized    Code = 3
	ErrDBNotUnique     Code = 105
	ErrRecordNotFound  Code = 106
	ErrJwtTokenInvalid Code = 202
	ErrQiniuUpload     Code = 301
	ErrNameUsed        Code = 1001
	ErrPasswordWrong   Code = 1002
	ErrUserNotExist    Code = 1003
	ErrNoPermission    Code = 1004
)

var codeMsg = map[Code]string{
	ErrServer:          "服务器错误",
	ErrBadRequest:      "错误的请求",
	ErrUnauthorized:    "未认证的请求",
	ErrDBNotUnique:     "唯一键冲突",
	ErrRecordNotFound:  "未找到记录",
	ErrJwtTokenInvalid: "无效Token",
	ErrQiniuUpload:     "文件上传错误",
	ErrNameUsed:        "名字已存在",
	ErrPasswordWrong:   "密码错误",
	ErrUserNotExist:    "用户名不存在",
	ErrNoPermission:    "无权限",
}

func (c Code) Error() string {
	if _, ok := codeMsg[c]; !ok {
		return "未知错误"
	}
	return codeMsg[c]
}

func (c Code) String() string {
	return c.Error()
}

func (c Code) With(err error) string {
	if err != nil {
		return c.Error() + ":" + err.Error()
	}
	return c.Error()
}
