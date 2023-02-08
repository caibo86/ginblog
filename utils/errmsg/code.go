package errmsg

type Code int

const (
	Success           Code = 0
	ErrServer         Code = 1
	ErrBadRequest     Code = 2
	ErrDBSelect       Code = 101
	ErrDBInsert       Code = 102
	ErrDBUpdate       Code = 103
	ErrDBDelete       Code = 104
	ErrDBNotUnique    Code = 105
	ErrRecordNotFound Code = 106
	ErrUsernameUsed   Code = 1001
	ErrPasswordWrong  Code = 1002
	ErrUserNotExist   Code = 1003
	ErrTokenExist     Code = 1004
	ErrTokenExpired   Code = 1005
	ErrTokenWrong     Code = 1006
	ErrTokenTypeWrong Code = 1007
)

var codeMsg = map[Code]string{
	Success:           "OK",
	ErrServer:         "服务器错误",
	ErrBadRequest:     "错误的请求",
	ErrDBSelect:       "数据库查询错误",
	ErrDBInsert:       "数据库插入错误",
	ErrDBUpdate:       "数据库更新错误",
	ErrDBDelete:       "数据库删除错误",
	ErrDBNotUnique:    "唯一键冲突",
	ErrRecordNotFound: "未找到记录",
	ErrUsernameUsed:   "用户名已存在",
	ErrPasswordWrong:  "密码错误",
	ErrUserNotExist:   "用户名不存在",
	ErrTokenExist:     "TOKEN不存在",
	ErrTokenExpired:   "TOKEN已过期",
	ErrTokenWrong:     "TOKEN不正确",
	ErrTokenTypeWrong: "TOKEN格式错误",
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
	return c.Error() + ":" + err.Error()
}
