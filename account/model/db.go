package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

// SetUp mysql
func SetUp(conf MysqlConnectConf) {
	dbHandle, err := DBConnect(conf)
	if err != nil {
		log.Fatal("connect db failure " + err.Error())
		os.Exit(1)
	}
	db = dbHandle
}

// TearDown mysql
func TearDown() {
	db.Close()
}

// Migrate database
func Migrate() {
	db.AutoMigrate(&Account{})
}

// MysqlConnectConf define mysql connect conf
type MysqlConnectConf struct {
	Host   string
	Port   int
	User   string
	Passwd string
	DB     string
}

// DBConnect use connect mysql
func DBConnect(conf MysqlConnectConf) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Passwd,
		conf.Host,
		conf.Port,
		conf.DB,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.DB().SetConnMaxLifetime(5 * 30) // 注意与mysql参数wait_timeout有关
	return
}
