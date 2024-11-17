package global

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Mode          string // debug | release
		Port          string
		DbType        string // mysql | sqlite
		DbAutoMigrate bool   // 是否自动迁移数据库表结构
		DbLogMode     string // silent | error | warn | info
	}
	Log struct {
		Level     string // debug | info | warn | error
		Prefix    string
		Format    string // text | json
		Directory string
	}
	Mysql struct {
		Host     string // 服务器地址
		Port     string // 端口
		Config   string // 高级配置
		Dbname   string // 数据库名
		Username string // 数据库用户名
		Password string // 数据库密码
	}
	Redis struct {
		DB       int    // 指定 Redis 数据库
		Addr     string // 服务器地址:端口
		Password string // 密码
	}
	Email struct {
		From     string // 发件人 要发邮件的邮箱
		Host     string // 服务器地址, 例如 smtp.qq.com 前往要发邮件的邮箱查看其 smtp 协议
		Port     int    // 前往要发邮件的邮箱查看其 smtp 协议端口, 大多为 465
		SmtpPass string // 邮箱密钥 不是密码是开启smtp后给你的密钥
		SmtpUser string // 邮箱账号
	}
	JWT struct {
		Secret string
		Expire int64 // hour
		Issuer string
	}
}

var Conf *Config

func GetConfig() *Config {
	if Conf == nil {
		log.Panic("配置还没有初始化")
		return nil
	}
	return Conf
}

func InitConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE

	if err := v.ReadInConfig(); err != nil {
		panic("配置文件读取失败: " + err.Error())
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic("配置文件反序列化失败: " + err.Error())
	}

	log.Println("配置文件内容加载成功: ", path)
	return Conf
}
