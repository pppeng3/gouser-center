package db

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	base = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	//获取dbconf
	mysqlConf := getDbConf()
	//构造dsn
	dsn := fmt.Sprintf(base, mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname)
	//连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "gormv2_",
			SingularTable: true,
		}}) //关闭默认事务提高性能
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	logrus.Info("mysql 连接成功")
	db.AutoMigrate()
}

func getDbConf() mysqlConf {
	viper.SetConfigName("db_conf")
	viper.AddConfigPath("./conf")
	if os.Getenv("ENV") == "dev" {
		logrus.Infoln("running in dev environment...")
		viper.AddConfigPath(os.Getenv("CODE") + "/gouser-center/config")
	}
	if err = viper.ReadInConfig(); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper readInconfig error")
	}
	var dbconf conf
	if err = viper.Unmarshal(&dbconf); err != nil {
		logrus.Error(errors.WithStack(err))
		panic("viper Unmarshal error")
	}
	return dbconf.Mysql
}

type conf struct {
	Mysql mysqlConf `json:"mysql"`
}

type mysqlConf struct {
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}
