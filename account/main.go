package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cyub/go-kit-demo/account/model"
	"github.com/cyub/go-kit-demo/account/router"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

const (
	appName      = "account-service"
	appVersion   = "0.1"
	appBuildTime = "2020-02-18"
)

func init() {
	var (
		consulHost  = flag.String("consul.host", "", "consul ip address")
		consulPort  = flag.String("consul.port", "", "consul port")
		serviceHost = flag.String("service.host", "8500", "service ip address")
		servicePort = flag.String("service.port", "8888", "service port")
		dbHost      = flag.String("mysql.host", "", "mysql ip address")
		dbPort      = flag.Int("mysql.port", 3360, "mysql port")
		dbUser      = flag.String("mysql.user", "", "mysql user")
		dbPasswd    = flag.String("mysql.passwd", "", "mysql password")
		dbDB        = flag.String("mysql.db", "", "mysql database")
	)
	flag.Parse()

	viper.Set("appName", appName)
	viper.Set("appVersion", appVersion)
	viper.Set("appBuildTime", appBuildTime)
	viper.Set("consulHost", *consulHost)
	viper.Set("consulPort", *consulPort)
	viper.Set("serviceHost", *serviceHost)
	viper.Set("servicePort", *servicePort)
	viper.Set("dbHost", *dbHost)
	viper.Set("dbPort", *dbPort)
	viper.Set("dbUser", *dbUser)
	viper.Set("dbPasswd", *dbPasswd)
	viper.Set("dbDB", *dbDB)
}

func main() {
	// 连接数据库
	dbConf := model.MysqlConnectConf{
		Host:   viper.GetString("dbHost"),
		Port:   viper.GetInt("dbPort"),
		User:   viper.GetString("dbUser"),
		Passwd: viper.GetString("dbPasswd"),
		DB:     viper.GetString("dbDB"),
	}
	model.SetUp(dbConf)
	// 迁移数据
	model.Migrate()
	// 关闭数据库连接
	defer model.TearDown()

	// 启动http服务
	errChan := make(chan error)
	go func() {
		r := router.NewRouter()
		fmt.Printf("Http server listen at:%s", viper.GetString("servicePort"))
		errChan <- http.ListenAndServe(":"+viper.GetString("servicePort"), r)
	}()

	err := <-errChan
	log.Fatal("http server crash " + err.Error())
}
