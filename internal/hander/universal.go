package hander

import (
	"blog/internal/cache"
	"blog/internal/global"

	"blog/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Universal struct{}

type SendCodeReq struct {
	Email string `json:"email"  binding:"required"`
}

// TODO 创建模版 区分发送邮箱类型
// 发送验证码
func (*Universal) SendCode(c *gin.Context) {
	var req SendCodeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, global.ErrRequest, err)
		return
	}
	validateCode := utils.GetCode()
	content := fmt.Sprintf(`
		<div style="text-align:center"> 
			<div style="padding: 8px 40px 8px 50px;">
            	<p>
					您本次的验证码为
					<p style="font-size:75px;font-weight:blod;"> %s </p>
					为了保证账号安全，验证码有效期为 %d 分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用~
				</p>
       	 	</div>
			<div>
            	<p>发送邮件专用邮箱，请勿回复。</p>
        	</div>
		</div>
	`, validateCode, 5)

	// 将验证码存储到 Redis 中
	cache.Set(GetRDB(c), global.KEY_SEND_CODE+req.Email, validateCode, time.Duration(5)*time.Minute)

	if err := utils.Email(req.Email, "博客注册验证码", content); err != nil {
		ReturnError(c, global.ErrSendEmail, err)
	}
	ReturnSuccess(c, nil)
}
