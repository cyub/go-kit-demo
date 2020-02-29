package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/cyub/go-kit-demo/account/model"
	"github.com/cyub/go-kit-demo/account/router"
	"github.com/cyub/go-kit-demo/account/service"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	appName      = "account-service"
	appVersion   = "0.1"
	appBuildTime = "2020-02-18"
)

func init() {
	var (
		consulHost = flag.String("consul.host", "consul", "consul ip address")
		consulPort = flag.Int("consul.port", 8500, "consul port")
		appHost    = flag.String("app.host", "8500", "app ip address")
		appPort    = flag.Int("app.port", 8888, "app port")
		appEnv     = flag.String("app.env", "dev", "app env")
	)
	flag.Parse()

	conf := service.NewConfig(
		*consulHost,
		*consulPort,
		appName,
		*appEnv,
		5)
	// 存储变量
	conf.Set("appName", appName)
	conf.Set("appVersion", appVersion)
	conf.Set("appBuildTime", appBuildTime)
	conf.Set("appHost", *appHost)
	conf.Set("appPort", *appPort)
	conf.Set("appEnv", *appEnv)
	conf.Set("consulHost", *consulHost)
	conf.Set("consulPort", *consulPort)
	// 加载配置
	conf.Load()

	// 日志配置
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

func main() {
	// 连接数据库
	dbConf := model.MysqlConnectConf{
		Host:   service.C.GetString("MYSQL_HOST"),
		Port:   service.C.GetInt("MYSQL_PORT"),
		User:   service.C.GetString("MYSQL_USER"),
		Passwd: service.C.GetString("MYSQL_PASSWD"),
		DB:     service.C.GetString("MYSQL_DB"),
	}
	model.SetUp(dbConf)
	// 迁移数据
	model.Migrate()
	// 关闭数据库连接
	defer model.TearDown()

	// 启动http服务
	errChan := make(chan error)
	go func() {
		r := router.New()
		r.Use(service.WithMetric(r))
		log.Info("Http server listen at:" + service.C.GetString("appPort"))
		errChan <- http.ListenAndServe(":"+service.C.GetString("appPort"), r)
	}()

	err := <-errChan
	log.Fatal("http server crash " + err.Error())
}
