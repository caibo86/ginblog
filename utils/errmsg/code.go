package errmsg

type Code int

const (
	OK                 Code = 0
	ErrServer          Code = 1
	ErrBadRequest      Code = 2
	ErrUnauthorized    Code = 3
	ErrDBSelect        Code = 101
	ErrDBInsert        Code = 102
	ErrDBUpdate        Code = 103
	ErrDBDelete        Code = 104
	ErrDBNotUnique     Code = 105
	ErrRecordNotFound  Code = 106
	ErrJwtSign         Code = 201
	ErrJwtTokenInvalid Code = 202
	ErrUsernameUsed    Code = 1001
	ErrPasswordWrong   Code = 1002
	ErrUserNotExist    Code = 1003
	ErrNoPermission    Code = 1004
)

var codeMsg = map[Code]string{
	OK:                 "OK",
	ErrServer:          "服务器错误",
	ErrBadRequest:      "错误的请求",
	ErrUnauthorized:    "未认证的请求",
	ErrDBSelect:        "数据库查询错误",
	ErrDBInsert:        "数据库插入错误",
	ErrDBUpdate:        "数据库更新错误",
	ErrDBDelete:        "数据库删除错误",
	ErrDBNotUnique:     "唯一键冲突",
	ErrRecordNotFound:  "未找到记录",
	ErrJwtSign:         "Jwt签名错误",
	ErrJwtTokenInvalid: "无效Token",
	ErrUsernameUsed:    "用户名已存在",
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
