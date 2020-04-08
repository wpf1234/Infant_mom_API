package base

import (
	"app/m/model"
	"app/m/utils"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const(
	defaultPath = "app"
)

var (
	m sync.RWMutex
	Key = []byte("infant*mom2020v1")
	mc model.MysqlConf
	DB *gorm.DB

	LC model.LogConf
	RC model.RedisConf
)

func Init(){
	m.Lock()
	defer m.Unlock()

	err:=config.Load(file.NewSource(
		file.WithPath("./conf/application.yml"),
	))

	if err!=nil{
		log.Error("加载配置文件失败: ",err)
		return
	}

	if err:=config.Get(defaultPath,"log").Scan(&LC);err!=nil{
		log.Error("Log 配置读取失败: ",err)
		return
	}
	log.Info("Log 配置读取成功!")
	utils.LoggerToFile(LC.LogPath,LC.LogFile)

	if err := config.Get(defaultPath, "mysql").Scan(&mc); err != nil {
		log.Error("Mysql 配置文件读取失败: ", err)
		return
	}
	log.Info("读取 Mysql 配置成功!")

	str, err := base64.StdEncoding.DecodeString(mc.Password)
	if err != nil {
		log.Error("Base64 decode failed: ", err)
		return
	}
	pwd, err := utils.AesDecrypt(str, Key)
	if err != nil {
		log.Error("ASE decrypt failed: ", err)
		return
	}
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true",
		mc.User, string(pwd), mc.Host, mc.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Error("Open mysql failed: ", err)
		return
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(10)
	DB.DB().SetConnMaxLifetime(10 * time.Second)

	if err := DB.DB().Ping(); err != nil {
		log.Error("连接数据库失败: ", err)
		return
	}
	log.Info("数据库连接成功!")

	if err := config.Get(defaultPath, "redis").Scan(&RC); err != nil {
		log.Error("Redis 配置文件读取失败: ", err)
		return
	}
	log.Info("读取 Redis 配置成功!")

	redisStr, err := base64.StdEncoding.DecodeString(RC.Password)
	if err != nil {
		log.Error("Base64 decode failed: ", err)
		return
	}
	redisPwd, err := utils.AesDecrypt(redisStr, Key)
	if err != nil {
		log.Error("ASE decrypt failed: ", err)
		return
	}
	err=utils.NewClient(RC.Host, string(redisPwd), RC.Port, RC.DB, RC.Pool)
	if err != nil {
		log.Error("创建 redis client 失败: ", err)
		return
	}

	log.Info("链接 redis 成功!")
}
