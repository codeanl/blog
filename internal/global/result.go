package global

var (
	_code = map[int]struct{}{}
	_msg  = make(map[int]string)
)

type Result struct {
	Code int
	Msg  string
}

func RegisterErrorCode(code int, msg string) Result {
	if _, ok := _code[code]; ok {
		panic("code has been registered")
	}
	if msg == "" {
		panic("msg cannot be empty")
	}

	_code[code] = struct{}{}
	_msg[code] = msg

	return Result{
		Code: code,
		Msg:  msg,
	}
}

var (
	OkReresult = RegisterErrorCode(200, "ok")
	FailResult = RegisterErrorCode(500, "fail")

	ErrRequest    = RegisterErrorCode(9001, "请求参数格式错误")
	ErrDbOp       = RegisterErrorCode(9002, "数据库操作异常")
	ErrBcryptHash = RegisterErrorCode(9003, "bcrypt加密异常")
	ErrSendEmail  = RegisterErrorCode(9004, "邮箱发送失败")

	ErrUserExist   = RegisterErrorCode(1001, "该用户名已经被注册")
	ErrUserNoExist = RegisterErrorCode(1004, "该用户不存在")
	ErrEmailExist  = RegisterErrorCode(1002, "该邮箱已经被注册")
	ErrPassword    = RegisterErrorCode(1003, "密码错误")
	ErrTokenCreate = RegisterErrorCode(1005, "token生成失败")
)
