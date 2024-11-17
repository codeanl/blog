package utils

import (
	"blog/internal/global"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"golang.org/x/exp/rand"
)

// 生成验证码
func GetCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(uint64(time.Now().UnixNano()))).Int31n(1000000))
}

//TODO 换一个包

// Email 发送方法
// to: 以英文逗号 , 分隔的字符串 如 "a@qq.com,b@qq.com"
func Email(to, subject, body string) error {
	return send(strings.Split(to, ","), subject, body)

}

// to: 目标数组
// subject: 邮件标题
// body: 邮件内容 (HTML)

func send(to []string, subject string, body string) error {
	conf := global.Conf
	// 从配置文件中读取
	from := conf.Email.From
	nickname := conf.Email.SmtpUser
	secret := conf.Email.SmtpUser
	host := conf.Email.Host
	port := conf.Email.Port

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", nickname, from) // 发件人
	e.To = to                                       // 收件人
	e.Subject = subject                             // 主题
	e.HTML = []byte(body)                           // 内容

	var err error
	auth := smtp.PlainAuth("", from, secret, host)
	addr := fmt.Sprintf("%s:%d", host, port)
	err = e.Send(addr, auth)
	return err
}
