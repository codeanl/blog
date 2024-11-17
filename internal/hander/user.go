package hander

import (
	"blog/internal/global"
	"blog/internal/model"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type User struct{}

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=4,max=20"`
	// Email    string `json:"email"  binding:"required"`
	// Code     string `json:"code"  binding:"required"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginVO struct {
	model.User
	Token string `json:"token"`
}

// TODO
// 用户注册之后会发送邮件 用户需要点击邮件中的链接进行激活 激活之后才能登录
// 邮箱会发送到用户的邮箱中 不需要收到获取验证码
// 先暂时用这样注册
func (*User) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, global.ErrRequest, err)
		return
	}

	// 检查用户名是否存在
	user, err := model.GetUserInfoByUsername(GetDB(c), req.Username)
	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}
	if user != nil {
		ReturnError(c, global.ErrUserExist, err)
		return
	}

	// // 检查邮箱是否存在
	// isExist, err := model.CheckEmailExist(GetDB(c), req.Email)
	// if err != nil {
	// 	ReturnError(c, global.ErrDbOp, err)
	// 	return
	// }
	// if isExist {
	// 	ReturnError(c, global.ErrEmailExist, nil)
	// 	return
	// }
	// 邮箱验证

	password, err := utils.BcryptHash(req.Password)
	if err != nil {
		ReturnError(c, global.ErrBcryptHash, err)
		return
	}

	err = model.Register(GetDB(c), &model.User{
		Username: req.Username,
		Password: password,
		// Email:    req.Email,
		Nickname: req.Username,
		Avatar:   global.DEFAULT_AVATAR,
	})
	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
	}
	ReturnSuccess(c, nil)
}

func (u *User) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, global.ErrRequest, err)
		return
	}
	user, err := model.GetUserInfoByUsername(GetDB(c), req.Username)
	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}
	if user == nil {
		ReturnError(c, global.ErrUserNoExist, err)
		return
	}

	if !utils.BcryptCheck(req.Password, user.Password) {
		ReturnError(c, global.ErrPassword, nil)
		return
	}

	// 生成Token
	conf := global.Conf.JWT
	token, err := utils.GenToken(conf.Secret, conf.Issuer, int(conf.Expire), user.ID)
	if err != nil {
		ReturnError(c, global.ErrTokenCreate, err)
		return
	}

	ReturnSuccess(c, LoginVO{
		User:  *user,
		Token: token,
	})
}
