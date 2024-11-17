package internal

import (
	"blog/internal/global"
	"blog/internal/model"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDatabase(conf *global.Config) *gorm.DB {

	var db *gorm.DB
	var err error

	//日志类型
	var level logger.LogLevel
	switch conf.Server.DbLogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		fallthrough
	default:
		level = logger.Error
	}

	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		},
	}

	dbtype := conf.Server.DbType
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Dbname, conf.Mysql.Config)

	switch dbtype {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), config)
	default:
		log.Fatal("不支持的数据库类型: ", dbtype)
	}
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	log.Println("数据库连接成功", dbtype, dsn)

	if conf.Server.DbAutoMigrate {
		if err := model.MakeMigrate(db); err != nil {
			log.Fatal("数据库迁移失败", err)
		}
		log.Println("数据库自动迁移成功")
	}

	return db
}

func InitRedis(conf *global.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("连接redis失败")
	}
	log.Println("连接redis成功")
	return rdb
}
